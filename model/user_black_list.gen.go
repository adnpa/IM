// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserBlackList = "user_black_list"

// UserBlackList mapped from table <user_black_list>
type UserBlackList struct {
	OwnerID    string    `gorm:"column:owner_id;primaryKey" json:"owner_id"`
	BlockID    string    `gorm:"column:block_id;primaryKey" json:"block_id"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time"`
}

// TableName UserBlackList's table name
func (*UserBlackList) TableName() string {
	return TableNameUserBlackList
}
