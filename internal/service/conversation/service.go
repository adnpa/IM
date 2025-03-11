package conversation

import (
	"context"

	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type ConversationService struct{}

func (s *ConversationService) GetAllHistoryConv(id int64) ([]*Conversation, error) {
	var result []*Conversation
	cur, _ := mongodb.GetAll("conversation", bson.M{
		"$or": []bson.M{
			{"left_user_id": id},
			{"right_user_id": id},
		}})
	err := cur.All(context.Background(), result)
	return result, err
}

func (s *ConversationService) SyncHistoryConv() {

}

