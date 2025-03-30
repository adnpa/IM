package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/internal/model"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

func (ws *WSServer) SendMsg(_ context.Context, in *pb.SendMsgReq) (*pb.SendMsgResp, error) {
	logger.Info("send msg to user", zap.Any("user_id", in.UserId), zap.Any("cmd", in.Msg.Typ), zap.Any("msg", in.Msg))

	conn, ok := ws.GetWsConn(int64(in.UserId))

	if conn == nil || !ok {
		return &pb.SendMsgResp{Succ: false}, fmt.Errorf("send conn is null")
	}

	var sendMsg model.CommonMsg
	sendMsg.Cmd = model.MsgType(in.Msg.Typ)
	switch in.Msg.Typ {
	case int32(model.TypMsgAckFromServerForSender):
		sendMsg.AckMsg = model.AckMsg{Id: in.Msg.Id, Seq: int32(in.Msg.Seq)}
	case int32(model.TypMsgAckFromServerForRecver):
		sendMsg.AckMsg = model.AckMsg{Id: in.Msg.Id, Seq: int32(in.Msg.Seq)}
	case int32(model.TypSingle):
		cmsg := model.Message{}
		copier.Copy(&cmsg, in.Msg)
		sendMsg.ChatMsg = cmsg
	case int32(model.TypGroup):
		cmsg := model.Message{}
		copier.Copy(&cmsg, in.Msg)
		sendMsg.ChatMsg = cmsg
	default:
		logger.Info("msg typ not recognized")
	}
	data, err := json.Marshal(sendMsg)
	if err != nil {
		logger.Error("marshal", zap.Error(err))
		return &pb.SendMsgResp{Succ: false}, err
	}
	logger.Info("end of send msg to user", zap.Any("user_id", in.UserId), zap.Any("resp msg", sendMsg))
	err = ws.writeMsg(conn, websocket.TextMessage, data)
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
