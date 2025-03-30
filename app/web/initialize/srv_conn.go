package initialize

import (
	"fmt"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/web/global"
	"github.com/adnpa/IM/pkg/common/discovery"
	"github.com/hashicorp/consul/api"
)

func InitSrvConn() {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	consulCli, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	userConn, err := discovery.GetGrpcConn(consulCli, "user-srv")
	if err != nil {
		panic(err)
	}
	global.UserCli = pb.NewUserClient(userConn)

	friendConn, err := discovery.GetGrpcConn(consulCli, "friend-srv")
	if err != nil {
		panic(err)
	}
	global.FriendCli = pb.NewFriendClient(friendConn)

	ossConn, err := discovery.GetGrpcConn(consulCli, "oss-srv")
	if err != nil {
		panic(err)
	}
	global.OssCli = pb.NewOSSClient(ossConn)

	// groupConn, err := discovery.GetGrpcConn(consulCli, "group-srv")
	// if err != nil {
	// 	panic(err)
	// }
	// global.GroupCli = pb.NewGroupClient(groupConn)

}
