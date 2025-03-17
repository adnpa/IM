package main

import (
	"flag"

	"github.com/adnpa/IM/internal/service/chat"
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

	chat.MyServer = &chat.WSServer{}
	wsServerPort := flag.Int("port", 10001, "get ginServerPort from cmd,default 10000 as port")
	flag.Parse()
	chat.MyServer.Init(*wsServerPort)
	chat.MyServer.Run()
}
