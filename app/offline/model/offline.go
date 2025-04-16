package model

import "github.com/adnpa/IM/internal/model"

type Inbox struct {
	Id    int32
	Inbox []model.ChatMessage
}
