package main

import (
	"flag"
	"github.com/adnpa/IM/internal/rpc/user"
	"github.com/adnpa/IM/pkg/common/config"
)

func main() {
	rpcPort := flag.Int("port", config.Config.RpcPort.UserPort[0], "RpcToken default listen port 10800")
	flag.Parse()
	rpcServer := user.NewUserServer(*rpcPort)
	rpcServer.Run()
}
