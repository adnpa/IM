package chat

import (
	"strconv"
	"time"

	"github.com/adnpa/IM/internal/model"
	"github.com/adnpa/IM/internal/service/offline"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"log"
	"net/http"
	"sync"
)

var (
	MyServer *WSServer
	rwLock   *sync.RWMutex
)

func init() {
	rwLock = &sync.RWMutex{}
}

type WsConn struct {
	*websocket.Conn
	w *sync.Mutex

	Seq int64
}

type WSServer struct {
	wsAddr     string
	maxConnNum int
	upgrader   *websocket.Upgrader
	mapConnUid map[*WsConn]string
	mapUidConn map[string]*WsConn
}

func (ws *WSServer) Init(wsPort int) {
	ip := utils.ServerIP
	ws.wsAddr = ip + ":" + utils.IntToString(wsPort)
	ws.maxConnNum = config.Config.LongConnSvr.WebsocketMaxConnNum
	ws.mapConnUid = make(map[*WsConn]string)
	ws.mapUidConn = make(map[string]*WsConn)
	ws.upgrader = &websocket.Upgrader{
		HandshakeTimeout: time.Duration(config.Config.LongConnSvr.WebsocketTimeOut) * time.Second,
		ReadBufferSize:   config.Config.LongConnSvr.WebsocketMaxMsgLen,
		CheckOrigin:      func(r *http.Request) bool { return true },
	}
}

func (ws *WSServer) Run() {
	log.Printf("Listening and serving ws on %s", ws.wsAddr)
	http.HandleFunc("/ws", ws.wsHandler)
	err := http.ListenAndServe(ws.wsAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (ws *WSServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	if ws.headerCheck(w, r) {
		query := r.URL.Query()
		// 升级为ws连接
		conn, err := ws.upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		} else {
			SendId := query.Get("uid")
			newConn := &WsConn{conn, new(sync.Mutex), 1}
			ws.AddWsConn(SendId, newConn)
			numId, _ := strconv.ParseInt(SendId, 10, 64)
			logger.Infof("user connect chat server", "id", numId)
			// 全量离线消息推送
			msgs := offline.GetOfflineMsg(numId)
			logger.Info("Start to sync msgs", zap.Any("msgs:", msgs))
			ws.SendMsg(newConn, model.CommonMsg{Cmd: model.TypSyncMsg, Msgs: msgs})
			logger.Infof("Sync msgs succ, end Login=================================")
			go ws.readMsg(newConn)
		}
	} else {
	}
}

// 验证登录
func (ws *WSServer) headerCheck(w http.ResponseWriter, r *http.Request) bool {
	// status := http.StatusUnauthorized
	// query := r.URL.Query()
	// if len(query["token"]) != 0 && len(query["sendID"]) != 0 && len(query["platformID"]) != 0 {
	// 	if !utils.VerifyToken(query["token"][0], query["sendID"][0]) {
	// 		w.Header().Set("Sec-Websocket-Version", "13")
	// 		http.Error(w, http.StatusText(status), status)
	// 		return false
	// 	} else {
	// 		return true
	// 	}
	// } else {
	// 	w.Header().Set("Sec-Websocket-Version", "13")
	// 	http.Error(w, http.StatusText(status), status)
	// 	return false
	// }
	return true
}

// -------------------------------------------------------

// 分布式支持
// 基于udp的分布式引用
