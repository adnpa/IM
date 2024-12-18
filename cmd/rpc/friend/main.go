package main

import (
	"flag"
	"github.com/adnpa/IM/internal/rpc/friend"
)

func main() {
	rpcPort := flag.Int("port", 10012, "RpcToken default listen port 10800")
	flag.Parse()
	rpcServer := friend.NewFriendServer(*rpcPort)
	rpcServer.Run()
}
