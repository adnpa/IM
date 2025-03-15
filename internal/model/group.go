package model

type GroupInfo struct {
	Gid      int64   `json:"gid,omitempty"`
	Name     string  `json:"name,omitempty"`
	Owner    int64   `json:"owner,omitempty"`
	Managers []int64 `json:"managers,omitempty"`
}

type GroupApply struct {
	GroupId     int64
	ApplyUserId int64
	// Flag        ApplyFlag //状态
}

// gid-[]Groupmember
type GroupMember struct {
	GroupId int64
	UserId  int64
	LastAck int64
}

type GroupMessages struct{}
