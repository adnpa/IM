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
		SendMq(data)
		logger.Info("msg", zap.Any("msg", data))
	}
}

// msg *pb.ChatMsg
func SendMq(body interface{}) error {
	return global.Producer.Send("msg", body)
}
