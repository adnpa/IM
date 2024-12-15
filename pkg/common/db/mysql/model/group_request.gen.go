// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGroupRequest = "group_request"

// GroupRequest mapped from table <group_request>
type GroupRequest struct {
	ID               int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GroupID          string    `gorm:"column:group_id;not null" json:"group_id"`
	FromUserID       string    `gorm:"column:from_user_id;not null" json:"from_user_id"`
	ToUserID         string    `gorm:"column:to_user_id;not null" json:"to_user_id"`
	Flag             int32     `gorm:"column:flag;not null" json:"flag"`
	ReqMsg           string    `gorm:"column:req_msg" json:"req_msg"`
	HandledMsg       string    `gorm:"column:handled_msg" json:"handled_msg"`
	CreateTime       time.Time `gorm:"column:create_time;not null" json:"create_time"`
	FromUserNickname string    `gorm:"column:from_user_nickname" json:"from_user_nickname"`
	ToUserNickname   string    `gorm:"column:to_user_nickname" json:"to_user_nickname"`
	FromUserFaceURL  string    `gorm:"column:from_user_face_url" json:"from_user_face_url"`
	ToUserFaceURL    string    `gorm:"column:to_user_face_url" json:"to_user_face_url"`
	HandledUser      string    `gorm:"column:handled_user" json:"handled_user"`
}

// TableName GroupRequest's table name
func (*GroupRequest) TableName() string {
	return TableNameGroupRequest
}