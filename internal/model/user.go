package model

import "time"

type Sex int

const (
	SEX_UNKNOW Sex = 0
	SEX_MEN
	SEX_WOMEN
)

type User struct {
	Id         int64
	Mobile     string
	Passwd     string
	Avatar     string
	Sex        Sex
	Nickname   string
	Salt       string
	OnlineTime int64 //上线时间
	Token      string
	Memo       string
	Createat   time.Time
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
