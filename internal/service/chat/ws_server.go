package chat

import (
	"strconv"
	"time"

	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/gorilla/websocket"

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
}

type WSServer struct {
	wsAddr     string
	maxConnNum int
	upgrader   *websocket.Upgrader
	connUserM  map[*WsConn]string
	userConnM  map[string]*WsConn
}

func (ws *WSServer) Init(wsPort int) {
	ip := utils.ServerIP
	ws.wsAddr = ip + ":" + utils.IntToString(wsPort)
	ws.maxConnNum = config.Config.LongConnSvr.WebsocketMaxConnNum
	ws.connUserM = make(map[*WsConn]string)
	ws.userConnM = make(map[string]*WsConn)
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
			newConn := &WsConn{conn, new(sync.Mutex)}
			ws.AddWsConn(SendId, newConn)
			numId, _ := strconv.ParseInt(SendId, 10, 64)
			logger.Infof("user connect chat server", "id", numId)
			msgs, _ := GetAllMsg(numId)
			logger.Infof("Start to sync msgs")
			ws.SendMsg(newConn, CommonMsg{Cmd: TypSyncMsg, Msgs: msgs})
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
// 连接管理
// -------------------------------------------------------
func (ws *WSServer) GetWsConn(uid int64) *WsConn {
	rwLock.RLock()
	defer rwLock.RUnlock()

	strUid := strconv.FormatInt(uid, 10)
	if conn, ok := ws.userConnM[strUid]; ok {
		return conn
	}
	return nil
}

func (ws *WSServer) AddWsConn(id string, c *WsConn) {
	rwLock.Lock()
	defer rwLock.Unlock()

	if oldConn, ok := ws.userConnM[id]; ok {
		err := oldConn.Close()
		delete(ws.connUserM, c)
		if err != nil {
			log.Println("close old conn error:", err)
		}
	} else {
		log.Println("first login", id)
	}

	ws.connUserM[c] = id
	ws.userConnM[id] = c
	logger.Infof("connect succ", "online users", ws.connUserM)
}

func (ws *WSServer) DelUserConn(conn *WsConn) {
	rwLock.Lock()
	defer rwLock.Unlock()

	if uid, ok := ws.connUserM[conn]; ok {
		if _, ok = ws.userConnM[uid]; ok {
			delete(ws.userConnM, uid)
		}
		delete(ws.connUserM, conn)
	}
	err := conn.Close()
	if err != nil {
		log.Println("close conn error:", err)
	}
}

func (ws *WSServer) GetUid(conn *WsConn) string {
	rwLock.RLock()
	defer rwLock.RUnlock()

	if conn, ok := ws.connUserM[conn]; ok {
		return conn
	}
	return ""
}

// -------------------------------------------------------

// 分布式支持
// 基于udp的分布式引用
