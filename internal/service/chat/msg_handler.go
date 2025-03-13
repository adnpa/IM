package chat

import (
	"encoding/json"
	"log"
	"time"

	"github.com/adnpa/IM/internal/service/group"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/logger"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// -------------------------------------------------------
// 消息管理
// -------------------------------------------------------

// -------------------------------------------------------
// 读协程 分发消息
// -------------------------------------------------------

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
		ws.handleMsg(conn, data)
	}
}

func (ws *WSServer) handleMsg(conn *WsConn, data []byte) {
	msg := &Message{}
	err := json.Unmarshal(data, msg)
	id := ws.GetUid(conn)
	logger.Infof("=================msg handle start====================")
	logger.Info("1 recv message", zap.String("source uid", id), zap.Any("msg", msg))

	if err != nil {
		logger.Error("unmarshal", zap.Error(err))
		return
	}

	switch msg.Cmd {
	case TypMsgAckFromClient:
		TransferQueue.PopMsg()
	case TypOfflineAck:
		HistoryMsgQueue.PopAllMsg(msg.From)
	case TypHeartbelt:
		ws.handleHeartbelt(conn)
	case TypSingle:
		if msg.To == msg.From {
			logger.Infof("message error to == from")
			return
		}
		msg.Id = utils.NowMilliSecond()
		TransferQueue.Product(msg)
		StoreMessage(msg)
		ws.SendMsg(conn, &CommonMsg{Cmd: TypMsgAckFromServer, Single: msg})
		time.Sleep(1 * time.Second)
	case TypGroup:
		if msg.To == msg.From {
			logger.Infof("message error to == from")
			return
		}
		StoreMessage(msg)
		msg.Id = utils.NowMilliSecond()
		users := group.GetAllGrouUser(msg.To)
		logger.Infof("send to all group members", "members", users)
		// 对群里所有用户 复制一条消息到队列
		for _, u := range users {
			tmp := msg
			tmp.RecverId = u.UserId
			TransferQueue.Product(tmp)
		}
		ws.SendMsg(conn, &CommonMsg{Cmd: TypMsgAckFromServer, Single: msg})
		time.Sleep(1 * time.Second)

	default:
		logger.Infof("unknown cmd")
	}
}

// todo online service
func (ws *WSServer) handleHeartbelt(conn *WsConn) {
	// uid := ws.GetUid(conn)
	// user := &user.User{}
	// user.OnlineTime = utils.NowMilliSecond()
	// mongodb.GetDecode("user", uid, user)
	// mongodb.Update("user", user)
	// ws.writeMsg(conn, websocket.PongMessage, []byte("pong"))
}

// -------------------------------------------------------
// 发送消息
// -------------------------------------------------------
// 写协程
// func (ws *WSServer) recvproc() {
// 	for {

// 	}
// }

func (ws *WSServer) SendMsg(conn *WsConn, reply interface{}) error {
	// todo 从map中获取map也要用锁
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

// func (ws *WSServer) getSeqReq(conn *WsConn, r *Req) {

// }

//func (ws *WServer) getSeqResp(conn *WsConn, m *Req, pb *pbChat.GetMaxAndMinSeqResp) {
//}

// func (ws *WSServer) sendErrMsg(conn *WsConn, errCode int32, errMsg string, reqIdentifier int32, msgIncr string, operationID string) {
// 	reply := Resp{
// 		ReqIdentifier: reqIdentifier,
// 		MsgIncr:       msgIncr,
// 		OperationID:   operationID,
// 		ErrCode:       errCode,
// 		ErrMsg:        errMsg,
// 	}
// 	ws.sendMsg(conn, reply)
// }
