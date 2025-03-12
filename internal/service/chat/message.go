package chat

import (
	"context"

	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type MsgType int

const (
	TypMsgAckFromServer MsgType = iota
	TypMsgAckFromClient
	TypOfflineMsg
	TypOfflineAck
	TypHeartbelt
	TypGroup
	TypSingle
	TypSyncMsg
)

type MediaType int

const (
	TextMessage MediaType = 1
)

type CommonMsg struct {
	Cmd    MsgType
	Single *Message
	Msgs   []*Message
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
	RecverId int64
}

// 弃用，和Message合并
// pk channel+msgId
// type GroupMessage struct {
// 	ChannerlId int64
// 	MsgId      int64
// 	From       int64
// 	Content    string
// }

func GetAllMsg(id int64) ([]*Message, error) {
	var result []*Message
	cur, _ := mongodb.GetAll("message", bson.M{
		"$or": []bson.M{
			{"from": id},
			{"to": id},
		}})
	err := cur.All(context.Background(), &result)
	return result, err
}
