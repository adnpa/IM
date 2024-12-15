package mongodb

import (
	"context"
	"github.com/adnpa/IM/pkg/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Conn struct {
	Client *mongo.Client
}

func GetConn() (*mongo.Client, error) {
	uri := "mongodb://localhost:27017"
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
}

func (c *Conn) Close() {
	err := c.Client.Disconnect(context.TODO())
	if err != nil {
		return
	}
}

func SaveUserChatMongo2(uid string, sendTime int64, m *pb.MsgDataToDB) error {

	return nil
}
