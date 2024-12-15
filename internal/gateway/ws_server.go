package gateway

import (
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/gorilla/websocket"
	"time"

	"log"
	"net/http"
	"sync"
)

var (
	rwLock *sync.RWMutex
)

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
	http.HandleFunc("/ws", ws.wsHandler)
	err := http.ListenAndServe(ws.wsAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (ws *WSServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	if ws.headerCheck(w, r) {
		query := r.URL.Query()
		conn, err := ws.upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		} else {
			//todo
			//SendID := query["sendID"][0] + " " + utils.PlatformIDToName(int32(utils.StringToInt64(query["platformID"][0])))
			SendId := query["sendID"][0]
			newConn := &WsConn{conn, new(sync.Mutex)}
			ws.addWsConn(SendId, newConn)
			go ws.readMsg(newConn)
		}
	}
}

// 验证登录
func (ws *WSServer) headerCheck(w http.ResponseWriter, r *http.Request) bool {
	status := http.StatusUnauthorized
	query := r.URL.Query()
	if len(query["token"]) != 0 && len(query["sendID"]) != 0 && len(query["platformID"]) != 0 {
		if !utils.VerifyToken(query["token"][0], query["sendID"][0]) {
			w.Header().Set("Sec-Websocket-Version", "13")
			http.Error(w, http.StatusText(status), status)
			return false
		} else {
			return true
		}
	} else {
		w.Header().Set("Sec-Websocket-Version", "13")
		http.Error(w, http.StatusText(status), status)
		return false
	}
}

//func (ws *WSServer) getWsConn(uid string) *WsConn {
//	rwLock.RLock()
//	defer rwLock.RUnlock()
//	if conn, ok := ws.userConnM[uid]; ok {
//		return conn
//	}
//	return nil
//}

func (ws *WSServer) getWsConn(uid, platform string) *WsConn {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if conn, ok := ws.userConnM[uid]; ok {
		return conn
	}
	return nil
}

func (ws *WSServer) addWsConn(id string, c *WsConn) {
	rwLock.Lock()
	defer rwLock.Unlock()

	//重复登录
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
}

func (ws *WSServer) delUserConn(conn *WsConn) {
	rwLock.Lock()
	defer rwLock.Unlock()
	if uid, ok := ws.connUserM[conn]; ok {
		if _, ok = ws.userConnM[uid]; ok {
			delete(ws.userConnM, uid)
		} else {
		}
		delete(ws.connUserM, conn)
	}
	err := conn.Close()
	if err != nil {
		log.Println("close conn error:", err)
	}
}

func (ws *WSServer) readMsg(conn *WsConn) {
	for {
		messageType, data, err := conn.ReadMessage()
		if messageType == websocket.PingMessage {
			log.Println("this is a  pingMessage")
		}
		if err != nil {
			ws.delConn(conn)
			return
		}

		ws.handleMsg(conn, data)
	}
}

func (ws *WSServer) writeMsg(conn *WsConn, msgType int, msg []byte) error {
	conn.w.Lock()
	defer conn.w.Unlock()
	return conn.WriteMessage(msgType, msg)
}

func (ws *WSServer) delConn(conn *WsConn) {
	rwLock.Lock()
	defer rwLock.Unlock()

	if uid, ok := ws.connUserM[conn]; ok {
		if _, ok := ws.userConnM[uid]; ok {
			delete(ws.userConnM, uid)
		}
		delete(ws.connUserM, conn)
	}

	err := conn.Close()
	if err != nil {
		log.Println("close conn error:", err)
	}
}
