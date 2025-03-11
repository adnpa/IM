package chat

import (
	"time"

	"github.com/adnpa/IM/internal/service/user"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/adnpa/IM/pkg/common/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

// 暂存消息

var TransferQueue *MessageQueue

// var HiQueue *MessageQueue

type MessageQueue struct {
	q *Deque[*Message]
}

func (mq *MessageQueue) Product(msg *Message) {
	mq.q.PushBack(msg)
	logger.Info("recvMs", zap.Any("msg", mq.q))

}

func (mq *MessageQueue) TrySendMsg() {
	// logger.Info("trysend")

	if mq.q.IsEmpty() {
		return
	}
	msg, _ := mq.q.PeekFront()
	user := &user.User{}
	mongodb.GetDecode("user", bson.M{"id": msg.To}, user)
	conn := MyServer.GetWsConn(msg.To)
	if IsOnline(user) {
		logger.Info("to User is online, send", zap.Any("msg", msg))
		MyServer.SendMsg(conn, CommonMsg{
			Cmd:    TypSingle,
			Single: msg,
		})
	} else {
		logger.Info("to User is not online,store", zap.Any("msg", msg))
		His.PutMsg(msg)
	}
}

func IsOnline(u *user.User) bool {
	return utils.NowMilliSecond()-u.OnlineTime > 5*time.Second.Milliseconds()
}

func (mq *MessageQueue) PopMsg() {
	mq.q.PopFront()
}

func Init() {
	TransferQueue = &MessageQueue{
		// MsgCh: make(chan *Message, 10),
		NewDeque[*Message](100),
	}

	His = &HistoryMsg{make(map[int64]*Deque[*Message])}
}

func Run() {
	for {
		TransferQueue.TrySendMsg()
		// 很重要 实时性到底要多少
		// time.Sleep(100 * time.Millisecond)
		time.Sleep(1 * time.Second)
	}
}

// transfer service 负责暂存消息并执行转发&推送通知操作
// todo push sdk支持
