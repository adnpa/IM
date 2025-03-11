package main

import (
	"flag"

	"github.com/adnpa/IM/internal/service/chat"
)

func main() {
	// 依赖服务
	chat.Init()
	go chat.Run()

	chat.MyServer = &chat.WSServer{}
	wsServerPort := flag.Int("port", 10001, "get ginServerPort from cmd,default 10000 as port")
	flag.Parse()
	chat.MyServer.Init(*wsServerPort)
	chat.MyServer.Run()
}
