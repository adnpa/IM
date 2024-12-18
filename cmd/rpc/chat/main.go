package main

import (
	"flag"
	"github.com/adnpa/IM/internal/rpc/chat"
)

func main() {
	rpcPort := flag.Int("port", 10011, "RpcToken default listen port 10800")
	flag.Parse()
	rpcServer := chat.NewRpcChatServer(*rpcPort)
	rpcServer.Run()
}
