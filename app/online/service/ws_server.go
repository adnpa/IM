package service

import (
	"strconv"
	"time"

	"github.com/adnpa/IM/app/online/model"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"net/http"
	"sync"
)

var (
	MyServer *WSServer
)

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
	rwLock     *sync.RWMutex
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

func (ws *WSServer) Run() error {
	http.HandleFunc("/ws", ws.HandleConn)
	return http.ListenAndServe(ws.wsAddr, nil)
}

func (ws *WSServer) HandleConn(w http.ResponseWriter, r *http.Request) {
	if ws.headerCheck(w, r) {
		query := r.URL.Query()
		conn, err := ws.upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		} else {
			logger.Info("", zap.Any("", query), zap.Any("", conn))
			SendId := query.Get("uid")
			newConn := &WsConn{conn, new(sync.Mutex), 1}
			ws.AddWsConn(SendId, newConn)
			numId, _ := strconv.ParseInt(SendId, 10, 64)
			// todo 暂时采用全量离线消息推送 后续结合会话多次推送
			msgs, _ := ws.GetOfflineMsg(numId)
			ws.SendMsg(newConn, model.CommonMsg{Cmd: model.TypSyncMsg, Msgs: msgs})
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

func (ws *WSServer) GetOfflineMsg(uid int64) ([]model.Message, error) {
	// todo 离线服务获取消息
	return nil, nil
}

// -------------------------------------------------------

// 分布式支持
// 基于udp的分布式引用
