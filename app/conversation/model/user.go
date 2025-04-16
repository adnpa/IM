package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Mobile    string         `gorm:"column:mobile" json:"mobile"`
	Email     string         `gorm:"column:email" json:"email"`
	Passwd    string         `gorm:"column:passwd" json:"passwd"`
	Salt      []byte         `gorm:"column:salt" json:"salt"`
	Nickname  string         `gorm:"column:nickname" json:"nickname"`
	Avatar    string         `gorm:"column:avatar" json:"avatar"`
	Sex       int32          `gorm:"column:sex" json:"sex"`
	Birthday  time.Time      `gorm:"column:birthday" json:"birthday"`
	Memo      string         `gorm:"column:memo" json:"memo"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
