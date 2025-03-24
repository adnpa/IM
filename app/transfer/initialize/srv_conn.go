package initialize

import (
	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/transfer/global"
	"github.com/adnpa/IM/pkg/common/discovery"
	"github.com/hashicorp/consul/api"
)

func InitSrvConn() {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	presenceConn, err := discovery.GetGrpcConn(consulCli, "presence-srv")
	if err != nil {
		panic(err)
	}
	global.PresenceCli = pb.NewPresenceClient(presenceConn)

	offlineConn, err := discovery.GetGrpcConn(consulCli, "offline-srv")
	if err != nil {
		panic(err)
	}
	global.OffineCli = pb.NewOfflineClient(offlineConn)

	// groupConn, err := discovery.GetGrpcConn(consulCli, "group-srv")
	// if err != nil {
	// 	panic(err)
	// }
	// global.GroupCli = pb.NewGroupClient(groupConn)

}
