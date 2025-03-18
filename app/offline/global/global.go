package global

import (
	"github.com/adnpa/IM/app/offline/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB           *mongo.Client
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
