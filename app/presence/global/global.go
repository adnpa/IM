package global

import (
	"github.com/adnpa/IM/app/user/config"
	"github.com/adnpa/IM/pkg/common/mq/rabbitmq"
	"github.com/adnpa/IM/pkg/common/storage"
)

var (
	Producer     *rabbitmq.Producer
	RedisPool    storage.Pool
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
