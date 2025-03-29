package global

import (
	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/transfer/config"
)

var (
	PresenceCli  pb.PresenceClient
	OffineCli    pb.OfflineClient
	GroupCli     pb.GroupClient
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
