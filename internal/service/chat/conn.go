package chat

import (
	"log"
	"strconv"

	"github.com/adnpa/IM/pkg/common/logger"
)

// -------------------------------------------------------
// 连接管理
// -------------------------------------------------------

func (ws *WSServer) GetWsConn(uid int64) (*WsConn, bool) {
	rwLock.RLock()
	defer rwLock.RUnlock()

	strUid := strconv.FormatInt(uid, 10)
	if conn, ok := ws.mapUidConn[strUid]; ok {
		return conn, true
	}
	return nil, false
}

func (ws *WSServer) AddWsConn(id string, c *WsConn) {
	rwLock.Lock()
	defer rwLock.Unlock()

	if oldConn, ok := ws.mapUidConn[id]; ok {
		err := oldConn.Close()
		delete(ws.mapConnUid, c)
		if err != nil {
			log.Println("close old conn error:", err)
		}
	} else {
		log.Println("first login", id)
	}

	ws.mapConnUid[c] = id
	ws.mapUidConn[id] = c
	logger.Infof("connect succ", "online users", ws.mapConnUid)
}

func (ws *WSServer) DelUserConn(conn *WsConn) {
	rwLock.Lock()
	defer rwLock.Unlock()

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
	rwLock.RLock()
	defer rwLock.RUnlock()

	if conn, ok := ws.mapConnUid[conn]; ok {
		return conn
	}
	return ""
}
