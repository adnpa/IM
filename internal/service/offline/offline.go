package offline

import (
	"github.com/adnpa/IM/internal/model"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/adnpa/IM/pkg/common/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type OfflineMsg struct {
	Uid  int64                   `json:"uid,omitempty"`
	Msgs map[int64]model.Message `json:"msgs,omitempty"`
}

func (m *OfflineMsg) PutMsg(msg model.Message) {
	m.Msgs[msg.Id] = msg
}

func (m *OfflineMsg) RemoveMsg(msgId int64) {
	delete(m.Msgs, msgId)
}

func (m *OfflineMsg) BatchRemoveMsg(msgIds []int64) {
	for _, id := range msgIds {
		m.RemoveMsg(id)
	}
}

func (m *OfflineMsg) GetMsg(id int64) (model.Message, bool) {
	msg, exist := m.Msgs[id]
	return msg, exist
}

func (m *OfflineMsg) GetMsgs() []model.Message {
	values := make([]model.Message, 0, len(m.Msgs))
	for _, value := range m.Msgs {
		values = append(values, value)
	}
	return values
}

func GetOfflineMsg(uid int64) []model.Message {
	oldMsg := &OfflineMsg{
		Uid:  uid,
		Msgs: make(map[int64]model.Message),
	}
	mongodb.GetDecode("offline", bson.M{"uid": uid}, oldMsg)
	return oldMsg.GetMsgs()
}

func Put(uid int64, msg model.Message) {
	logger.Info("put", zap.Any("", uid), zap.Any("", msg))
	oldMsg := &OfflineMsg{
		Uid:  uid,
		Msgs: make(map[int64]model.Message),
	}
	mongodb.GetDecode("offline", bson.M{"uid": uid}, oldMsg)
	if len(oldMsg.Msgs) == 0 {
		oldMsg.PutMsg(msg)
		mongodb.Insert("offline", oldMsg)
		return
	} else {
		oldMsg.PutMsg(msg)
		update := bson.M{
			"$set": bson.M{
				"msgs": oldMsg.Msgs,
			},
		}
		mongodb.Update("offline", bson.M{"uid": uid}, update)
	}
}

func Remove(uid int64, msgId int64) {
	oldMsg := &OfflineMsg{
		Uid:  uid,
		Msgs: make(map[int64]model.Message),
	}
	mongodb.GetDecode("offline", bson.M{"uid": uid}, oldMsg)
	oldMsg.RemoveMsg(msgId)
	update := bson.M{
		"$set": bson.M{
			"msgs": oldMsg.Msgs,
		},
	}
	mongodb.Update("offline", bson.M{"uid": uid}, update)
}

func Clear(uid int64) {
	mongodb.Delete("offline", bson.M{"uid": uid})
}
