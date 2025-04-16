package model

import "encoding/json"

type MsgType int

const (
	CmdChat MsgType = iota + 1
	CmdPullOfflineMsgs
	CmdAck
)

const (
	TypMsgAckFromServerForSender MsgType = iota + 1
	// TypMsgAckFromClient
	TypOfflineMsg
	TypOfflineAck
	// TypHeartbelt
	// TypGroup
	// TypSingle
	// TypSyncMsg
	// TypMsgAckFromServerForRecver
	// TypPullOfflineMsg
	// TypPullOfflineMsgResp
	// TypAckPullOffline

)

const (
	ChatTypeSingle = iota + 1
	ChatTypeGroup
)

type CommonMsg struct {
	Cmd    MsgType         `json:"cmd,omitempty"`    // 消息类型标识
	Data   json.RawMessage `json:"data,omitempty"`   // 根据Cmd解析具体类型
	Seq    int64           `json:"seq,omitempty"`    //序列号
	Source int64           `json:"source,omitempty"` //来源
}

type AckMsg struct {
	Id  int64 `json:"id,omitempty"`
	Seq int32 `json:"seq,omitempty"`
}

// --聊天消息
type MediaType int

const (
	TextMessage MediaType = iota + 1
	ImageMessage
	AudioMessage
	VideoMessage
)

type ChatMessage struct {
	Id           int64     `json:"id,omitempty" form:"id"`           // 消息ID
	Type         MsgType   `json:"type,omitempty" form:"type"`       // 消息类型 群聊/单聊
	From         int64     `json:"from,omitempty" form:"from"`       // 发送者ID
	To           int64     `json:"to,omitempty" form:"to"`           // 接收者ID或群组ID
	Media        MediaType `json:"media,omitempty" form:"media"`     // 媒体类型
	Content      string    `json:"content,omitempty" form:"content"` // 消息内容
	Pic          string    `json:"pic,omitempty" form:"pic"`         // 缩略图URL
	Url          string    `json:"url,omitempty" form:"url"`         // 服务URL
	Memo         string    `json:"memo,omitempty" form:"memo"`       // 备注
	Amount       int       `json:"amount,omitempty" form:"amount"`   // 数字相关，如语音长度等
	Seq          int64     `json:"seq,omitempty"`                    //序列号
	Conversation int64     `json:"conversation,omitempty"`           //所属会话
}

// 群消息 弃用，和Message合并
// pk channel+msgId
// type GroupMessage struct {
// 	ChannerlId int64
// 	MsgId      int64
// 	From       int64
// 	Content    string
// }

//-- 拉离线消息

type PullOfflineMsgReq struct{}

type PullOfflineMsgResp struct {
	Msgs []ChatMessage `json:"msgs,omitempty"`
}


// -- Ack

type Ack struct{
	Seq int64
}