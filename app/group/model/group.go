package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameGroup = "group"

// Group 群聊表
type Group struct {
	GroupID     int64          `gorm:"column:group_id;primaryKey;autoIncrement:true;comment:群聊ID，主键，自增" json:"group_id"` // 群聊ID，主键，自增
	GroupName   string         `gorm:"column:group_name;not null;comment:群聊名称" json:"group_name"`                        // 群聊名称
	CreatorID   int32          `gorm:"column:creator_id;not null;comment:创建者用户ID" json:"creator_id"`                     // 创建者用户ID
	AvatarURL   string         `gorm:"column:avatar_url;comment:群聊头像URL" json:"avatar_url"`                              // 群聊头像URL
	Description string         `gorm:"column:description;comment:群聊描述" json:"description"`                               // 群聊描述
	MaxMembers  int32          `gorm:"column:max_members;default:200;comment:群聊最大成员数" json:"max_members"`                // 群聊最大成员数
	Status      int32          `gorm:"column:status;comment:群聊状态（0:正常，1:解散）" json:"status"`                              // 群聊状态（0:正常，1:解散）
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`       // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`       // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                 // 删除时间
}

// TableName Group's table name
func (*Group) TableName() string {
	return TableNameGroup
}
