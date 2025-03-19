package global

import (
	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/user/config"
	"github.com/adnpa/IM/pkg/common/storage"
)

var (
	RedisPool    storage.Pool
	PresenceCli  pb.PresenceClient
	OffineCli    pb.OfflineClient
	GroupCli     pb.GroupClient
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
