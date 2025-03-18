package chat

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"strconv"

// 	"github.com/adnpa/IM/internal/model"
// 	"github.com/adnpa/IM/internal/service/group"
// 	"github.com/adnpa/IM/internal/service/offline"
// 	"github.com/adnpa/IM/internal/utils"
// 	"github.com/adnpa/IM/pkg/common/logger"

// 	"github.com/gorilla/websocket"
// 	"go.uber.org/zap"
// )

// // 消息管理

// func (ws *WSServer) readMsg(conn *WsConn) {
// 	for {
// 		messageType, data, err := conn.ReadMessage()
// 		if err != nil {
// 			ws.DelUserConn(conn)
// 			return
// 		}
// 		if messageType == websocket.PingMessage {
// 			log.Println("ping msg from client")
// 		}
// 		ws.handleMsg(conn, data)
// 	}
// }

// // func SendMq(msg *model.Message) error {
// // 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// // 	if err != nil {
// // 		logger.Panic("Failed to connect to RabbitMQ", zap.Error(err))
// // 	}
// // 	defer conn.Close()
// // 	return rabbitmq.NewProducer(conn, "").Send("msg", msg)
// // }

// func (ws *WSServer) handleMsg(conn *WsConn, data []byte) {
// 	msg := &model.Message{}
// 	err := json.Unmarshal(data, msg)
// 	id := ws.GetUid(conn)
// 	logger.Infof("=================msg handle start====================")
// 	logger.Info("recv message", zap.String("source uid", id), zap.Any("msg", msg))

// 	if err != nil {
// 		logger.Error("unmarshal", zap.Error(err))
// 		return
// 	}

// 	switch msg.Cmd {
// 	case model.TypMsgAckFromClient:
// 		uid, _ := strconv.ParseInt(ws.GetUid(conn), 10, 64)
// 		offline.Remove(uid, msg.Id)
// 		ws.SendMsg(conn, &model.CommonMsg{Cmd: model.TypMsgAckFromServerForRecver, Single: msg})
// 	case model.TypOfflineAck:
// 		offline.Clear(msg.From)
// 	case model.TypHeartbelt:
// 		// ws.handleHeartbelt(conn)
// 	case model.TypSingle:
// 		ws.handleSingle(conn, msg)
// 	case model.TypGroup:
// 		ws.handleGroup(conn, msg)
// 	default:
// 		logger.Infof("unknown cmd")
// 	}
// }

// func (ws *WSServer) handleSingle(conn *WsConn, msg *model.Message) {
// 	if conn.Seq > msg.Seq || ((msg.Cmd == model.TypSingle || msg.Cmd == model.TypGroup) && msg.To == msg.From) {
// 		logger.Infof("message error to == from")
// 		return
// 	}

// 	msg.Id = utils.NowMilliSecond()
// 	if recverConn, ok := ws.GetWsConn(msg.To); ok {
// 		logger.Info("to User is online, send", zap.Any("msg", msg))
// 		ws.SendMsg(recverConn, model.CommonMsg{
// 			Cmd:    msg.Cmd,
// 			Single: msg,
// 		})
// 	}
// 	offline.Put(msg.To, *msg)
// 	// 历史消息
// 	StoreMessage(msg)
// 	ws.SendMsg(conn, &model.CommonMsg{Cmd: model.TypMsgAckFromServerForSender, Single: msg})
// 	conn.Seq++
// }

// func (ws *WSServer) handleGroup(conn *WsConn, msg *model.Message) {
// 	if conn.Seq > msg.Seq || ((msg.Cmd == model.TypSingle || msg.Cmd == model.TypGroup) && msg.To == msg.From) {
// 		logger.Infof("message error to == from")
// 		return
// 	}

// 	msg.Id = utils.NowMilliSecond()
// 	users := group.GetAllGrouUser(msg.To)
// 	logger.Infof("send to all group members", "members", users)
// 	// 对群里所有用户 复制一条消息到队列
// 	for _, u := range users {
// 		tmp := msg
// 		if u.UserId == msg.From {
// 			continue
// 		}
// 		// 这里应该不能改成功
// 		tmp.RecverId = u.UserId

// 		// 存在长连接说明在线
// 		if recverConn, ok := ws.GetWsConn(u.UserId); ok {
// 			logger.Info("to User is online, push", zap.Any("send to", u.UserId), zap.Any("msg", msg))
// 			ws.SendMsg(recverConn, model.CommonMsg{
// 				Cmd:    msg.Cmd,
// 				Single: tmp,
// 			})
// 		}
// 		offline.Put(u.UserId, *msg)
// 	}
// 	StoreMessage(msg)
// 	ws.SendMsg(conn, &model.CommonMsg{Cmd: model.TypMsgAckFromServerForSender, Single: msg})
// 	conn.Seq++
// }

// // todo online service
// func (ws *WSServer) handleHeartbelt(conn *WsConn) {
// 	// uid := ws.GetUid(conn)
// 	// user := &user.User{}
// 	// user.OnlineTime = utils.NowMilliSecond()
// 	// mongodb.GetDecode("user", uid, user)
// 	// mongodb.Update("user", user)
// 	// ws.writeMsg(conn, websocket.Pongmodel.Message, []byte("pong"))
// }

// // -------------------------------------------------------
// // 发送消息
// // -------------------------------------------------------
// // 写协程
// // func (ws *WSServer) recvproc() {
// // 	for {

// // 	}
// // }

// func (ws *WSServer) SendMsg(conn *WsConn, reply interface{}) error {
// 	// todo 从map中获取map也要用锁
// 	if conn == nil {
// 		logger.Warn("send conn is null")
// 		return fmt.Errorf("send conn is null")
// 	}
// 	data, err := json.Marshal(reply)
// 	if err != nil {
// 		return err
// 	}
// 	err = ws.writeMsg(conn, websocket.BinaryMessage, data)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return nil
// }

// func (ws *WSServer) writeMsg(conn *WsConn, msgType int, msg []byte) error {
// 	conn.w.Lock()
// 	defer conn.w.Unlock()

// 	return conn.WriteMessage(msgType, msg)
// }

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

// func (ws *WSServer) IsOnline(uid int64) bool {
// 	logger.Info("is online ", zap.Any("", ws.GetWsConn(uid)), zap.Any("", ws.mapUidConn))
// 	return ws.GetWsConn(uid) != nil
// }
