package model

import "time"

type Friend struct {
	OwnerID    int64     `gorm:"column:owner_id;primaryKey" json:"owner_id,omitempty"`
	FriendID   int64     `gorm:"column:friend_id;primaryKey" json:"friend_id,omitempty"`
	Comment    string    `gorm:"column:comment" json:"comment,omitempty"`
	Group      int64     `json:"group,omitempty"` // 分组标识 用于前端分组管理
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time,omitempty"`
}

// type Friend struct {
// 	Id       int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
// 	Ownerid  int64     `xorm:"bigint(20)" form:"ownerid" json:"ownerid"` // 记录是谁的
// 	Dstobj   int64     `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`   // 对端信息
// 	Cate     int       `xorm:"int(11)" form:"cate" json:"cate"`          // 什么类型
// 	Memo     string    `xorm:"varchar(120)" form:"memo" json:"memo"`     // 备注
// 	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"` // 创建时间
// }
