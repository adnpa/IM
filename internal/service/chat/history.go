package chat

// histroy 历史消息(离线消息)
var HistoryMsgQueue *HistoryMsg

type HistoryMsg struct {
	history map[int64]*Deque[*Message]
}

func (mq *HistoryMsg) PullUsrOffLineMsg(uid int64) {
	queue := mq.history[uid]
	conn := MyServer.GetWsConn(uid)

	MyServer.SendMsg(conn, &CommonMsg{
		Cmd:  TypOfflineMsg,
		Msgs: queue.buffer,
	})
}

func (mq *HistoryMsg) PutMsg(msg *Message) {
	if q, ok := mq.history[msg.To]; ok {
		q.PushBack(msg)
	} else {
		mq.history[msg.To] = NewDeque[*Message](100)
	}
}

func (mq *HistoryMsg) PopAllMsg(uid int64) {
	delete(mq.history, uid)
}
