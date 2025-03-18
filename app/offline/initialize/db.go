package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/adnpa/IM/app/offline/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() {
	c := global.ServerConfig.MongoInfo
	var err error
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		c.Name,
		c.Password,
		c.Host,
		c.Port)
	global.DB, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}
