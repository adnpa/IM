package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameGroupMember = "group_member"

// GroupMember 群聊成员表
type GroupMember struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键，自增" json:"id"`                   // 主键，自增
	GroupID   int64          `gorm:"column:group_id;not null;comment:群聊ID，外键" json:"group_id"`                          // 群聊ID，外键
	UserID    int32          `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                               // 用户ID
	JoinTime  time.Time      `gorm:"column:join_time;not null;default:CURRENT_TIMESTAMP;comment:加入时间" json:"join_time"` // 加入时间
	Role      int32          `gorm:"column:role;comment:成员角色（0:普通成员，1:管理员，2:群主）" json:"role"`                           // 成员角色（0:普通成员，1:管理员，2:群主）
	Status    int32          `gorm:"column:status;comment:成员状态（0:正常，1:已退出）" json:"status"`                              // 成员状态（0:正常，1:已退出）
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`        // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`        // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                  // 删除时间
}

// TableName GroupMember's table name
func (*GroupMember) TableName() string {
	return TableNameGroupMember
}
