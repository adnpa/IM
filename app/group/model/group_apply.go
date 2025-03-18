package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameGroupApply = "group_apply"

// GroupApply 群聊申请表
type GroupApply struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:申请ID，主键，自增" json:"id"`       // 申请ID，主键，自增
	GroupID     int64          `gorm:"column:group_id;not null;comment:群聊ID，外键" json:"group_id"`                   // 群聊ID，外键
	ApplicantID int32          `gorm:"column:applicant_id;not null;comment:申请人用户ID" json:"applicant_id"`           // 申请人用户ID
	Status      int32          `gorm:"column:status;comment:申请状态（0:待处理，1:已通过，2:已拒绝）" json:"status"`                // 申请状态（0:待处理，1:已通过，2:已拒绝）
	HandlerID   int32          `gorm:"column:handler_id;comment:处理人用户ID" json:"handler_id"`                        // 处理人用户ID
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
}

// TableName GroupApply's table name
func (*GroupApply) TableName() string {
	return TableNameGroupApply
}
