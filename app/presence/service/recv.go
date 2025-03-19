package service

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/adnpa/IM/app/presence/global"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func (ws *WSServer) readMsg(conn *WsConn) {
	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			ws.DelUserConn(conn)
			return
		}
		if messageType == websocket.PingMessage {
			log.Println("ping msg from client")
		}
		// ws.handleMsg(conn, data)
		SendMq(data)
		logger.Info("msg", zap.Any("msg", data))
	}
}

// msg *pb.ChatMsg
func SendMq(body interface{}) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Panic("Failed to connect to RabbitMQ", zap.Error(err))
	}
	defer conn.Close()
	return global.Producer.Send("msg", body)
}
