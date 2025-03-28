package main

import (
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
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

// import (
// 	"context"
// 	"flag"
// 	"log"

// 	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
// 	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
// )

// // 定义全局变量
// var (
// 	region     string // 存储区域
// 	bucketName string // 存储空间名称
// )

// // init函数用于初始化命令行参数
// func init() {
// 	flag.StringVar(&region, "region", "", "The region in which the bucket is located.")
// 	flag.StringVar(&bucketName, "bucket", "", "The name of the bucket.")
// }

func main() {
	initialize.InitConfig()
	initialize.InitAliOssCli()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// dbURL := os.Getenv("OSS_ACCESS_KEY_ID")
	// apiKey := os.Getenv("OSS_ACCESS_KEY_SECRET")
	// log.Println(dbURL)
	// log.Println(apiKey)
	// log.Println(os.Getwd())
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50057))
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
		GRPC:                           fmt.Sprintf("%s:%d", utils.ServerIP, 50057),
		Timeout:                        "3s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "10s",
	}

	serverId := uuid.NewString()
	registeration := api.AgentServiceRegistration{
		ID:      serverId,
		Address: utils.ServerIP,
		Port:    50057,
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
