package service

import (
	"log"

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
		logger.Info("msg", zap.Any("", data))
	}
}
