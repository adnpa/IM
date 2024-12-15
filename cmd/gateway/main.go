package main

import (
	"flag"
	"github.com/adnpa/IM/internal/gateway"
	"sync"
)

func main() {
	rpcPort := flag.Int("rpc_port", 10400, "rpc listening port")
	wsPort := flag.Int("ws_port", 17778, "ws listening port")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	gateway.Init(*rpcPort, *wsPort)
	gateway.Run()
	wg.Wait()
}
