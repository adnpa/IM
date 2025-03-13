package test

import (
	"testing"

	"github.com/adnpa/IM/internal/service/chat"
)

func TestMes(t *testing.T) {
	msgs, _ := chat.GetAllMsg(2)
	t.Log(msgs)
}
