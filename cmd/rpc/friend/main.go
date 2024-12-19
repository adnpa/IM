package main

import (
	"flag"
	"github.com/adnpa/IM/internal/rpc/friend"
	"github.com/adnpa/IM/pkg/common/config"
)

func main() {
	rpcPort := flag.Int("port", config.Config.RpcPort.FriendPort[0], "RpcToken default listen port 10800")
	flag.Parse()
	rpcServer := friend.NewFriendServer(*rpcPort)
	rpcServer.Run()
}
