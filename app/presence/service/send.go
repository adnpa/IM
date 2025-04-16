package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/internal/model"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func (ws *WSServer) SendMsg(_ context.Context, in *pb.SendMsgReq) (*pb.SendMsgResp, error) {
	logger.Info("send msg to user", zap.Any("user_id", in.UserId), zap.Any("msg", in))

	conn, ok := ws.GetWsConn(int64(in.UserId))

	if conn == nil || !ok {
		return &pb.SendMsgResp{Succ: false}, fmt.Errorf("send conn is null")
	}

	var sendMsg model.CommonMsg
	sendMsg.Cmd = model.MsgType(in.Type)
	sendMsg.Data = in.Body
	// switch sendMsg.Cmd {
	// // case model.TypMsgAckFromServerForSender:
	// // 	sendMsg.AckMsg = model.AckMsg{Id: in.Msg.Id, Seq: int32(in.Msg.Seq)}
	// // case model.TypMsgAckFromServerForRecver:
	// // 	sendMsg.AckMsg = model.AckMsg{Id: in.Msg.Id, Seq: int32(in.Msg.Seq)}
	// case model.CmdChat:
	// 	cmsg := model.ChatMessage{}
	// 	copier.Copy(&cmsg, in.Msg)
	// 	data, err := json.Marshal(cmsg)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	sendMsg.Data = data
	// case model.CmdPullOfflineMsgs:
	// 	var msgs []model.ChatMessage
	// 	for _, pbMsg := range in.Msgs {
	// 		msg := model.ChatMessage{}
	// 		copier.Copy(&msg, pbMsg)
	// 		msgs = append(msgs, msg)
	// 	}
	// 	data, err := json.Marshal(model.PullOfflineMsgResp{Msgs: msgs})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	sendMsg.Data = data
	// default:
	// 	logger.Info("msg typ not recognized")
	// }
	// logger.Info("end of send msg to user", zap.Any("user_id", in.UserId), zap.Any("resp msg", sendMsg))

	if err := ws.sendMsg(conn, sendMsg); err != nil {
		logger.Error("write msg", zap.Error(err))
		return &pb.SendMsgResp{Succ: false}, err
	}
	return &pb.SendMsgResp{Succ: true}, nil
}

// func (ws *WSServer) writeMsg(conn *WsConn, msgType int, msg []byte) error {
// 	conn.w.Lock()
// 	defer conn.w.Unlock()

// 	return conn.WriteMessage(msgType, msg)
// }

func (ws *WSServer) sendMsg(conn *WsConn, sendMsg interface{}) error {
	conn.w.Lock()
	defer conn.w.Unlock()

	data, err := json.Marshal(sendMsg)
	if err != nil {
		logger.Error("marshal", zap.Error(err))
		return err
	}
	return conn.WriteMessage(websocket.TextMessage, data)
}
