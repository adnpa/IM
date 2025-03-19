package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func (ws *WSServer) SendMsg(_ context.Context, in *pb.SendMsgReq) (*pb.SendMsgResp, error) {
	logger.Info("get send msg")
	conn, ok := ws.GetWsConn(int64(in.UserId))

	if conn == nil || !ok {
		return &pb.SendMsgResp{Succ: false}, fmt.Errorf("send conn is null")
	}
	data, err := json.Marshal(in.Msg)
	if err != nil {
		logger.Error("marshal", zap.Error(err))
		return &pb.SendMsgResp{Succ: false}, err
	}
	err = ws.writeMsg(conn, websocket.BinaryMessage, data)
	if err != nil {
		logger.Error("write msg", zap.Error(err))
		return &pb.SendMsgResp{Succ: false}, err
	}
	return &pb.SendMsgResp{Succ: true}, nil
}

func (ws *WSServer) writeMsg(conn *WsConn, msgType int, msg []byte) error {
	conn.w.Lock()
	defer conn.w.Unlock()

	return conn.WriteMessage(msgType, msg)
}

func (ws *WSServer) sendMsg(conn *WsConn, reply interface{}) error {
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
