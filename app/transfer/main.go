package main

import (
	"fmt"
	"log"

	"github.com/adnpa/IM/app/transfer/global"
	"github.com/adnpa/IM/app/transfer/initialize"
	"github.com/adnpa/IM/app/transfer/service"
	"github.com/adnpa/IM/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func main() {
	// 初始化
	initialize.InitConfig()
	logger.Info("", zap.Any("", global.ServerConfig.ConsulInfo))

	initialize.InitSrvConn()
	
	rabbitCfg := global.ServerConfig.RabbitMQInfo
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbitCfg.User, rabbitCfg.Password, rabbitCfg.Host, rabbitCfg.Port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := service.NewConsumer(conn)
	go c.Run()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}
	// 	debug.SetMemoryLimit(512 * 1024 * 1024)

	// 	// var port = flag.Int("port", 50055, "The server port")
	// 	// flag.Parse()
	// 	// ws服务
	// 	wsServer := &service.WSServer{}
	// 	wsServer.Init(utils.ServerIP, 50055, 50056)
	// 	//grpc服务
	// 	s := grpc.NewServer()
	// 	pb.RegisterPresenceServer(s, wsServer)
	// 	healthcheck := health.NewServer()
	// 	healthgrpc.RegisterHealthServer(s, healthcheck)

	// 	// 服务注册
	// 	cfg := api.DefaultConfig()
	// 	cfg.Address = "localhost:8500"
	// 	cli, err := api.NewClient(cfg)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	check := api.AgentServiceCheck{
	// 		GRPC:                           utils.ServerIP,
	// 		Timeout:                        "3s",
	// 		Interval:                       "10s",
	// 		DeregisterCriticalServiceAfter: "10s",
	// 	}

	// 	registeration := api.AgentServiceRegistration{
	// 		ID:      wsServer.ServerId(),
	// 		Address: utils.ServerIP,
	// 		Port:    wsServer.PortGrpc,
	// 		Name:    global.ServerConfig.Name,
	// 		// Tags:    tags,
	// 		Check: &check,
	// 	}

	// 	err = cli.Agent().ServiceRegister(&registeration)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// 开启服务
	// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", wsServer.PortGrpc))
	// 	http.HandleFunc("/ws", wsServer.HandleConn)
	// 	logger.Info("ws listen on", zap.String("addr", wsServer.Host), zap.Int("port", wsServer.PortWs))
	// 	logger.Info("grpc listen on", zap.String("addr", wsServer.Host), zap.Int("port", wsServer.PortGrpc))

	// 	go func() {
	// 		err = http.ListenAndServe(fmt.Sprintf("%s:%d", wsServer.Host, wsServer.PortWs), nil)
	// 		if err != nil {
	// 			panic("failed to start ws server:" + err.Error())
	// 		}
	// 	}()

	// 	go func() {
	// 		err = s.Serve(lis)
	// 		if err != nil {
	// 			panic("failed to start grpc:" + err.Error())
	// 		}
	// 	}()

	// 	//接收终止信号
	// 	quit := make(chan os.Signal, 1)
	// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 	<-quit
	// 	if err = cli.Agent().ServiceDeregister(wsServer.ServerId()); err != nil {
	// 		zap.S().Info("注销失败")
	// 	}
	// 	zap.S().Info("注销成功")

}
