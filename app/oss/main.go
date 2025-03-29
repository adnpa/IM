package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/oss/global"
	"github.com/adnpa/IM/app/oss/initialize"
	"github.com/adnpa/IM/app/oss/service"
	"github.com/adnpa/IM/internal/utils"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	initialize.InitConfig()
	initialize.InitAliOssCli()

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file", err)
	// }

	var port = flag.Int("port", 50057, "The server port")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
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
		GRPC:                           fmt.Sprintf("%s:%d", utils.ServerIP, *port),
		Timeout:                        "3s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "10s",
	}

	serverId := uuid.NewString()
	registeration := api.AgentServiceRegistration{
		ID:      serverId,
		Address: utils.ServerIP,
		Port:    *port,
		Name:    global.ServerConfig.Name,
		// Tags:    tags,
		Check: &check,
	}

	err = cli.Agent().ServiceRegister(&registeration)
	if err != nil {
		panic(err)
	}

	pb.RegisterOSSServer(s, &service.OssService{})
	go func() {
		err = s.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}

	}()
	log.Println("start srver")
	select {}
}
