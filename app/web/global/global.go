package global

import (
	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/web/config"
)

var (
	GroupCli     pb.GroupClient
	FriendCli    pb.FriendClient
	UserCli      pb.UserClient
	OssCli       pb.OSSClient
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
