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
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=admin&connectTimeoutMS=3000",
		c.User,
		c.Password,
		c.Host,
		c.Port)

	opts := options.Client().ApplyURI(uri).
		SetConnectTimeout(3 * time.Second).
		SetServerSelectionTimeout(3 * time.Second)
	global.DB, err = mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}

	if err = global.DB.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
}
