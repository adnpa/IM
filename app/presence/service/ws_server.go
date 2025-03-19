package service

import (
	"context"
	"strconv"
	"time"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/presence/global"
	"github.com/adnpa/IM/app/presence/model"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"net/http"
	"sync"
)

type Heartbeat struct{}

type WsConn struct {
	*websocket.Conn
	w *sync.Mutex

	Seq int64
}

// ws *WSServer pb.PresenceServer
type WSServer struct {
	pb.UnimplementedPresenceServer

	serverId string
	Host     string
	PortGrpc int
	PortWs   int

	maxConnNum int
	upgrader   *websocket.Upgrader
	mapConnUid map[*WsConn]string
	mapUidConn map[string]*WsConn
	rwLock     *sync.RWMutex
}

func (ws *WSServer) Init(host string, wsPort int, grpcPort int) {
	// ip := utils.ServerIP
	ws.Host = host
	ws.PortWs = wsPort
	ws.PortGrpc = grpcPort
	ws.maxConnNum = config.Config.LongConnSvr.WebsocketMaxConnNum
	ws.mapConnUid = make(map[*WsConn]string)
	ws.mapUidConn = make(map[string]*WsConn)
	ws.upgrader = &websocket.Upgrader{
		HandshakeTimeout: time.Duration(config.Config.LongConnSvr.WebsocketTimeOut) * time.Second,
		ReadBufferSize:   config.Config.LongConnSvr.WebsocketMaxMsgLen,
		CheckOrigin:      func(r *http.Request) bool { return true },
	}
	ws.rwLock = &sync.RWMutex{}
	ws.serverId = uuid.NewString()
}

func (ws *WSServer) ServerId() string {
	return ws.serverId
}

// func (ws *WSServer) Run() error {
// 	http.HandleFunc("/ws", ws.HandleConn)
// 	log.Printf("ws server listening at %s:%d", ws.host, ws.wsPort)
// 	s := grpc.NewServer()
// 	s.ser
// 	return http.ListenAndServe(fmt.Sprintf("%s:%d", ws.host, ws.wsPort), nil)
// }

func (ws *WSServer) HandleConn(w http.ResponseWriter, r *http.Request) {
	if ws.headerCheck(w, r) {
		logger.Info("user connect", zap.Any("url", r.URL))
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
			ws.sendMsg(newConn, model.CommonMsg{Cmd: model.TypSyncMsg, Msgs: msgs})
			go ws.readMsg(newConn)

			// redis记录用户连接的服务器 心跳刷新ttl
			conn, err := global.RedisPool.Get(context.Background())
			if err != nil {
				return
			}
			conn.SetNX(SendId, ws.serverId, 6*time.Minute)
		}
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
	return []model.Message{}, nil
}

// -------------------------------------------------------

// 分布式支持
// 基于udp的分布式引用
