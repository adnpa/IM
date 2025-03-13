package chat

import (
	"time"

	"github.com/adnpa/IM/internal/service/group"
	"github.com/adnpa/IM/internal/service/user"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/adnpa/IM/pkg/common/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

// transfer service 负责暂存消息并执行转发&推送通知操作
// todo push sdk支持

var TransferQueue *MessageQueue

func Init() {
	TransferQueue = &MessageQueue{
		NewDeque[*Message](100),
	}

	HistoryMsgQueue = &HistoryMsg{make(map[int64]*Deque[*Message])}
}

func Run() {
	for {
		TransferQueue.TransferMsg()
		// 实时性 轮询转发消息
		// time.Sleep(100 * time.Millisecond)
		time.Sleep(1 * time.Second)
	}
}

type MessageQueue struct {
	q *Deque[*Message]
}

func (mq *MessageQueue) Product(msg *Message) {
	mq.q.PushBack(msg)
	logger.Info("msg enqueue transfer queue")
}

func (mq *MessageQueue) TransferMsg() {
	if mq.q.IsEmpty() {
		return
	}
	msg, _ := mq.q.PeekFront()
	user := &user.User{}
	mongodb.GetDecode("user", bson.M{"id": msg.To}, user)
	var conn *WsConn
	if msg.Cmd == TypGroup {
		conn = MyServer.GetWsConn(msg.RecverId)
	} else {
		conn = MyServer.GetWsConn(msg.To)
	}
	if IsOnline(user) {
		logger.Info("to User is online, send", zap.Any("msg", msg))
		MyServer.SendMsg(conn, CommonMsg{
			Cmd:    msg.Cmd,
			Single: msg,
		})
	} else {
		logger.Info("to User is not online,store", zap.Any("msg", msg))
		HistoryMsgQueue.PutMsg(msg)
	}
}

func (mq *MessageQueue) PopMsg() {
	msg, _ := mq.q.PopFront()
	if msg.Cmd == TypGroup {
		gm := &group.GroupMember{}
		mongodb.GetDecode("group_member", bson.M{"uid": msg.RecverId}, gm)
		gm.LastAck = msg.Id
		mongodb.Update("group_member", gm)
	}
}

func IsOnline(u *user.User) bool {
	// todo 后续在线服务要改成心跳包判断
	// return utils.NowMilliSecond()-u.OnlineTime > 5*time.Second.Milliseconds()
	// return MyServer.GetWsConn(u.Id) != nil
	return true
}
