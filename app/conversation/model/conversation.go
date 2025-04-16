package model

import (
	"time"

	"github.com/adnpa/IM/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConversationType int

// 群消息

type Conversation struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type         ConversationType   `bson:"type" json:"type"`
	Participants []int32            `bson:"participants" json:"participants"` // 参与者ID列表
	LastMessage  *model.ChatMessage `bson:"last_message" json:"last_message"` // 最后一条消息快照
	UnreadCount  map[string]int     `bson:"unread_count" json:"unread_count"` // 用户ID -> 未读数
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	// 可扩展字段
	Extensions map[string]interface{} `bson:"extensions,omitempty" json:"extensions,omitempty"`
}

// type Conversation struct {
// 	UserId    int   `json:"id,omitempty"`
// 	Type      int   `json:"type,omitempty"`
// 	Target    int64 `json:"target,omitempty"`     //
// 	Line      int64 `json:"line,omitempty"`       //分线
// 	UnreadNum int   `json:"unread_num,omitempty"` //未读数
// }
