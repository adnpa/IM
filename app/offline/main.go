package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/offline/initialize"
	"github.com/adnpa/IM/app/offline/service"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// 初始化
	// initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	var port = flag.Int("port", 50051, "The server port")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &service.UserService{})
	log.Printf("server listening at %v", lis.Addr())
	// 支持grpc健康检查
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(s, healthcheck)

	// 服务注册
	cfg := api.DefaultConfig()
	cfg.Address = "localhost:8500"
	cli, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 注意,check service是检查连接,连接是什么就用哪种类型,这里是tcp连接
	check := api.AgentServiceCheck{
		GRPC:                           "localhost:50051",
		Timeout:                        "3s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "10s",
	}

	serverId := uuid.NewString()
	registeration := api.AgentServiceRegistration{
		ID:      serverId,
		Address: "192.168.151.66",
		Port:    50051,
		Name:    "user-srv",
		// Tags:    tags,
		Check: &check,
	}

	err = cli.Agent().ServiceRegister(&registeration)
	if err != nil {
		panic(err)
	}

	// 开启服务
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
	go func() {
		err = s.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = cli.Agent().ServiceDeregister(serverId); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")

}
