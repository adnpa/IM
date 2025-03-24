package model

type MsgType int

const (
	TypMsgAckFromServerForSender MsgType = iota + 1
	TypMsgAckFromClient
	TypOfflineMsg
	TypOfflineAck
	TypHeartbelt
	TypGroup
	TypSingle
	TypSyncMsg
	TypMsgAckFromServerForRecver
)

type MediaType int

const (
	TextMessage MediaType = 1
)

type CommonMsg struct {
	Cmd     MsgType `json:"cmd,omitempty"`
	ChatMsg Message `json:"chat_msg,omitempty"`
	AckMsg  AckMsg  `json:"ack_msg,omitempty"`
}

type AckMsg struct {
	Id  int64 `json:"id,omitempty"`
	Seq int32 `json:"seq,omitempty"`
}

// pk msgId
type Message struct {
	Id       int64   `json:"id,omitempty" form:"id"`           // 消息ID
	Cmd      MsgType `json:"cmd,omitempty" form:"cmd"`         // 消息类型
	From     int64   `json:"from,omitempty" form:"from"`       // 发送者ID
	To       int64   `json:"to,omitempty" form:"to"`           // 接收者ID或群组ID
	Media    int     `json:"media,omitempty" form:"media"`     // 媒体类型
	Content  string  `json:"content,omitempty" form:"content"` // 消息内容
	Pic      string  `json:"pic,omitempty" form:"pic"`         // 缩略图URL
	Url      string  `json:"url,omitempty" form:"url"`         // 服务URL
	Memo     string  `json:"memo,omitempty" form:"memo"`       // 备注
	Amount   int     `json:"amount,omitempty" form:"amount"`   // 数字相关，如语音长度等
	Seq      int64   `json:"seq,omitempty"`
	RecverId int64   `json:"recver_id,omitempty"`
}

// 弃用，和Message合并
// pk channel+msgId
// type GroupMessage struct {
// 	ChannerlId int64
// 	MsgId      int64
// 	From       int64
// 	Content    string
// }
