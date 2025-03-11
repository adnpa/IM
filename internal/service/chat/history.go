package chat

// histroy 历史消息(离线消息)
// 实现在线功能后再实现
var His *HistoryMsg

type HistoryMsg struct {
	// msg chan[*chat.Message]
	history map[int64]*Deque[*Message]
}

type OfflineMsg struct {
	Cmd  MsgType
	Msgs []*Message
}

func (mq *HistoryMsg) GetUsrOffLineMsg(uid int64) {
	queue := mq.history[uid]
	conn := MyServer.GetWsConn(uid)
	MyServer.SendMsg(conn, &OfflineMsg{
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
