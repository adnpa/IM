package gateway

import (
	"bytes"
	"encoding/gob"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/gorilla/websocket"
	"log"
)

func (ws *WSServer) handleMsg(conn *WsConn, data []byte) {
	b := bytes.NewBuffer(data)
	m := Req{}
	dec := gob.NewDecoder(b)
	err := dec.Decode(&m)
	if err != nil {
		log.Println(err)
	}

	switch m.ReqIdentifier {
	case constant.WSGetNewestSeq:
		ws.getSeqReq(conn, &m)
	default:
	}
}

func (ws *WSServer) getSeqReq(conn *WsConn, r *Req) {

}

//func (ws *WServer) getSeqResp(conn *WsConn, m *Req, pb *pbChat.GetMaxAndMinSeqResp) {
//}

func (ws *WSServer) sendMsg(conn *WsConn, reply interface{}) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(reply)
	if err != nil {
		log.Println(err)
	}

	err = ws.writeMsg(conn, websocket.BinaryMessage, b.Bytes())
	if err != nil {
		log.Println(err)
	}
}

func (ws *WSServer) sendErrMsg(conn *WsConn, errCode int32, errMsg string, reqIdentifier int32, msgIncr string, operationID string) {
	reply := Resp{
		ReqIdentifier: reqIdentifier,
		MsgIncr:       msgIncr,
		OperationID:   operationID,
		ErrCode:       errCode,
		ErrMsg:        errMsg,
	}
	ws.sendMsg(conn, reply)
}
