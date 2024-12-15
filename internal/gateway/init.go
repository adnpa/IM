package gateway

import (
	"sync"
)

var (
	ws     WSServer
	rpcSvr RpcRelayServer
)

func Init(rpcPort, wsPort int) {
	rwLock = new(sync.RWMutex)
	ws.Init(wsPort)
	rpcSvr.Init(rpcPort)
}

func Run() {
	go ws.Run()
	go rpcSvr.Run()
}
