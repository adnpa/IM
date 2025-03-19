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
	Memo      string         `gorm:"column:memo" json:"memo"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
