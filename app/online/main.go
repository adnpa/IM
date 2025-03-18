package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/adnpa/IM/app/online/initialize"
	"github.com/adnpa/IM/app/online/service"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

func main() {
	// errors.New()
	// 依赖服务
	// chat.Init()
	// go chat.Run()

	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// if err != nil {
	// 	logger.Panic("Failed to connect to RabbitMQ", zap.Error(err))
	// }
	// defer conn.Close()
	// go rabbitmq.NewConsumer(conn).Run()

	initialize.InitConfig()
	initialize.InitDB()

	var port = flag.Int("port", 50051, "The server port")
	flag.Parse()
	wsServer := &service.WSServer{}
	wsServer.Init(*port)

	// 服务注册
	cfg := api.DefaultConfig()
	cfg.Address = "localhost:8500"
	cli, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	check := api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("localhost:%d", *port),
		Timeout:                        "3s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "10s",
	}

	serverId := uuid.NewString()
	registeration := api.AgentServiceRegistration{
		ID:      serverId,
		Address: "192.168.151.66",
		Port:    50051,
		Name:    "ws-srv",
		// Tags:    tags,
		Check: &check,
	}

	err = cli.Agent().ServiceRegister(&registeration)
	if err != nil {
		panic(err)
	}

	// 开启服务
	go func() {
		err = wsServer.Run()
		if err != nil {
			panic("failed to start ws server:" + err.Error())
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
