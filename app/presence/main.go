package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/presence/global"
	"github.com/adnpa/IM/app/presence/initialize"
	"github.com/adnpa/IM/app/presence/service"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// 初始化
	initialize.InitConfig()
	initialize.InitDB()
	initialize.InitProducer()

	// var port = flag.Int("port", 50055, "The server port")
	// flag.Parse()
	// ws服务
	wsServer := &service.WSServer{}
	wsServer.Init(utils.ServerIP, 50055, 50056)
	//grpc服务
	s := grpc.NewServer()
	pb.RegisterPresenceServer(s, wsServer)
	// 支持grpc健康检查
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(s, healthcheck)

	// 服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	cli, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	check := api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", utils.ServerIP, wsServer.PortGrpc),
		Timeout:                        "3s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "10s",
	}

	registeration := api.AgentServiceRegistration{
		ID:      wsServer.ServerId(),
		Address: utils.ServerIP,
		Port:    wsServer.PortGrpc,
		Name:    global.ServerConfig.Name,
		// Tags:    tags,
		Check: &check,
	}

	err = cli.Agent().ServiceRegister(&registeration)
	if err != nil {
		panic(err)
	}

	// 开启服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", wsServer.PortGrpc))
	http.HandleFunc("/ws", wsServer.HandleConn)
	logger.Info("ws listen on", zap.String("addr", wsServer.Host), zap.Int("port", wsServer.PortWs))
	logger.Info("grpc listen on", zap.String("addr", wsServer.Host), zap.Int("port", wsServer.PortGrpc))

	go func() {
		err = http.ListenAndServe(fmt.Sprintf("%s:%d", wsServer.Host, wsServer.PortWs), nil)
		if err != nil {
			panic("failed to start ws server:" + err.Error())
		}
	}()

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
	if err = cli.Agent().ServiceDeregister(wsServer.ServerId()); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")

}
