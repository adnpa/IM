package service

import (
	"context"
	"time"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/presence/global"
	"github.com/adnpa/IM/internal/model"
	"github.com/adnpa/IM/internal/utils"
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
	query := r.URL.Query()
	token := query.Get("token")
	if user_id, ok := ws.tokenCheck(token); ok {
		logger.Info("user connect", zap.String("user_id", user_id), zap.Any("url", r.URL))
		conn, err := ws.upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Error("connect error", zap.Error(err))
			return
		} else {
			newConn := &WsConn{conn, new(sync.Mutex), 1}
			ws.AddWsConn(user_id, newConn)
			// numId, _ := strconv.ParseInt(user_id, 10, 64)
			// TODO: 暂时采用全量离线消息推送 后续结合会话多次推送
			// msgs, _ := ws.GetOfflineMsg(numId)
			// ws.sendMsg(newConn, model.CommonMsg{Cmd: model.TypSyncMsg, Msgs: msgs})
			go ws.readMsg(newConn)

			// redis记录用户连接的服务器 心跳刷新ttl
			conn, err := global.RedisPool.Get(context.Background())
			if err != nil {
				logger.Error("push redis fail", zap.Error(err))
				return
			}
			conn.SetNX(user_id, ws.serverId, 6*time.Minute)
		}
	}
}

// 验证登录
func (ws *WSServer) tokenCheck(token string) (string, bool) {
	// authHeader := r.Header.Get("authorization")
	// token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return "", false
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		return "", false
	}
	// intUid, _ := strconv.ParseInt(claims.UID, 10, 64)
	return claims.UID, true
}

func (ws *WSServer) GetOfflineMsg(uid int64) ([]model.Message, error) {
	// TODO: 离线服务获取消息
	return []model.Message{}, nil
}

// -------------------------------------------------------

// 分布式支持
// 基于udp的分布式引用
