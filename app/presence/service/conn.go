package service

import (
	"context"
	"log"
	"strconv"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/presence/global"
	"github.com/adnpa/IM/pkg/logger"
	"go.uber.org/zap"
)

func (ws *WSServer) IsOnline(_ context.Context, in *pb.IsOnlineReq) (*pb.IsOnlineResp, error) {
	cli, err := global.RedisPool.Get(context.Background())
	if err != nil {
		return nil, err
	}

	resp, err := cli.Get(strconv.Itoa(int(in.UserId)))
	if err != nil {
		return nil, err
	}

	return &pb.IsOnlineResp{IsOnline: resp != "", ServerId: resp}, nil
}

// -------------------------------------------------------
// 连接管理
// -------------------------------------------------------

func (ws *WSServer) GetWsConn(uid int64) (*WsConn, bool) {
	ws.rwLock.RLock()
	defer ws.rwLock.RUnlock()

	strUid := strconv.FormatInt(uid, 10)
	if conn, ok := ws.mapUidConn[strUid]; ok {
		return conn, true
	}
	return nil, false
}

func (ws *WSServer) AddWsConn(id string, c *WsConn) {
	ws.rwLock.Lock()
	defer ws.rwLock.Unlock()

	if oldConn, ok := ws.mapUidConn[id]; ok {
		err := oldConn.Close()
		delete(ws.mapConnUid, c)
		if err != nil {
			logger.Warn("close old conn error", zap.Error(err))
		}
	}

	ws.mapConnUid[c] = id
	ws.mapUidConn[id] = c
}

func (ws *WSServer) DelUserConn(conn *WsConn) {
	ws.rwLock.Lock()
	defer ws.rwLock.Unlock()

	if uid, ok := ws.mapConnUid[conn]; ok {
		if _, ok = ws.mapUidConn[uid]; ok {
			delete(ws.mapUidConn, uid)
		}
		delete(ws.mapConnUid, conn)
	}
	err := conn.Close()
	if err != nil {
		log.Println("close conn error:", err)
	}
}

func (ws *WSServer) GetUid(conn *WsConn) string {
	ws.rwLock.RLock()
	defer ws.rwLock.RUnlock()

	if conn, ok := ws.mapConnUid[conn]; ok {
		return conn
	}
	return ""
}
