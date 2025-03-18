package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFriendship = "friendships"

// Friendship 好友关系表
type Friendship struct {
	UserID    int32          `gorm:"column:user_id;primaryKey;comment:用户ID" json:"user_id"`                      // 用户ID
	FriendID  int32          `gorm:"column:friend_id;primaryKey;comment:好友ID" json:"friend_id"`                  // 好友ID
	Comment   string         `gorm:"column:comment;comment:备注" json:"comment"`                                   // 备注
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
}

// TableName Friendship's table name
func (*Friendship) TableName() string {
	return TableNameFriendship
}
