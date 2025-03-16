package model

import (
	"time"
)

type Sex int

const (
	SEX_UNKNOW Sex = 0
	SEX_MEN
	SEX_WOMEN
)

type User struct {
	Id         int64     `json:"id,omitempty"`
	Mobile     string    `json:"mobile,omitempty"`
	Email      string    `json:"email,omitempty"`
	Passwd     string    `json:"passwd,omitempty"`
	Salt       []byte    `json:"salt,omitempty"`
	Nickname   string    `json:"nickname,omitempty"`
	Avatar     string    `json:"avatar,omitempty"`
	Sex        Sex       `json:"sex,omitempty"`
	OnlineTime int64     `json:"online_time,omitempty"` //上线时间
	Token      string    `json:"token,omitempty"`
	Memo       string    `json:"memo,omitempty"`
	Createat   time.Time `json:"createat,omitempty"`
}

// type User struct {
// 	UID        string    `gorm:"column:uid;primaryKey" json:"uid"`
// 	Name       string    `gorm:"column:name" json:"name"`
// 	Icon       string    `gorm:"column:icon" json:"icon"`
// 	Gender     int32     `gorm:"column:gender" json:"gender"`
// 	Mobile     string    `gorm:"column:mobile" json:"mobile"`
// 	Birth      string    `gorm:"column:birth" json:"birth"`
// 	Email      string    `gorm:"column:email" json:"email"`
// 	Ex         string    `gorm:"column:ex" json:"ex"`
// 	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
// }
