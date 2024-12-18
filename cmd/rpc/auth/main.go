package main

import (
	"flag"
	"github.com/adnpa/IM/internal/rpc/auth"
)

func main() {
	rpcPort := flag.Int("port", 10010, "RpcToken default listen port 10800")
	flag.Parse()
	rpcServer := auth.NewAuthServer(*rpcPort)
	rpcServer.Run()
}
