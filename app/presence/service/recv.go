package service

import (
	"context"
	"log"

	"github.com/adnpa/IM/app/presence/global"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func (ws *WSServer) readMsg(conn *WsConn) {
	for {
		messageType, data, err := conn.ReadMessage()
		logger.Info("recv msg from user", zap.Any("msg", data))

		if err != nil {
			ws.DelUserConn(conn)
			redisConn, err := global.RedisPool.Get(context.Background())
			if err != nil {
				return
			}
			redisConn.Del(ws.GetUid(conn))
			return
		}
		if messageType == websocket.PingMessage {
			log.Println("ping msg from client")
		}
		// ws.handleMsg(conn, data)
		err = SendMq(data)
		if err != nil {
			logger.Error("publish msg to mq fail", zap.Error(err))
		}
	}
}

// msg *pb.ChatMsg
// func SendMq(body interface{}) error {
// 	return global.Producer.Send("im_message", body)
// }

func SendMq(data []byte) error {
	return global.Producer.Send("im_message", data)
}
