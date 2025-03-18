package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/adnpa/IM/pkg/logger"
	"github.com/gorilla/websocket"
)

func (ws *WSServer) SendMsg(conn *WsConn, reply interface{}) error {
	// todo 从map中获取map也要用锁
	if conn == nil {
		logger.Warn("send conn is null")
		return fmt.Errorf("send conn is null")
	}
	data, err := json.Marshal(reply)
	if err != nil {
		return err
	}
	err = ws.writeMsg(conn, websocket.BinaryMessage, data)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (ws *WSServer) writeMsg(conn *WsConn, msgType int, msg []byte) error {
	conn.w.Lock()
	defer conn.w.Unlock()

	return conn.WriteMessage(msgType, msg)
}
