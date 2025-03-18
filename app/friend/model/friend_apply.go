package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFriendApply = "friend_apply"

// FriendApply 好友申请表
type FriendApply struct {
	ID          int32          `gorm:"column:id;primaryKey;autoIncrement:true;comment:申请ID" json:"id"`             // 申请ID
	FromID      int32          `gorm:"column:from_id;not null;comment:申请者ID" json:"from_id"`                       // 申请者ID
	ToID        int32          `gorm:"column:to_id;not null;comment:被申请者ID" json:"to_id"`                          // 被申请者ID
	Status      int32          `gorm:"column:status;comment:申请状态" json:"status"`                                   // 申请状态
	ApplyReason string         `gorm:"column:apply_reason;comment:申请理由" json:"apply_reason"`                       // 申请理由
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
}

// TableName FriendApply's table name
func (*FriendApply) TableName() string {
	return TableNameFriendApply
}
