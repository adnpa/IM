package chat

import (
	"context"
	"github.com/adnpa/IM/common/constant"
	"github.com/adnpa/IM/pkg/pb"
)

type MsgCallBackReq struct {
	SendID       string `json:"sendID"`
	RecvID       string `json:"recvID"`
	Content      string `json:"content"`
	SendTime     int64  `json:"sendTime"`
	MsgFrom      int32  `json:"msgFrom"`
	ContentType  int32  `json:"contentType"`
	SessionType  int32  `json:"sessionType"`
	PlatformID   int32  `json:"senderPlatformID"`
	MsgID        string `json:"msgID"`
	IsOnlineOnly bool   `json:"isOnlineOnly"`
}
type MsgCallBackResp struct {
	ErrCode         int32  `json:"errCode"`
	ErrMsg          string `json:"errMsg"`
	ResponseErrCode int32  `json:"responseErrCode"`
	ResponseResult  struct {
		ModifiedMsg string `json:"modifiedMsg"`
		Ext         string `json:"ext"`
	}
}

func (rpc *RpcChatServer) SendMsg(_ context.Context, req *pb.SendMsgReq) (*pb.SendMsgResp, error) {
	reply := &pb.SendMsgResp{}

	switch req.MsgData.SessionType {
	case constant.SingleChatType:
		//canSend, err := callbackBeforeSendSingle(req)
		//if err != nil {
		//	return nil, err
		//}

		msgToMQ := pb.MsgDataToMQ{Token: req.Token, OperationID: req.OperationID, MsgData: req.MsgData}
		err := rpc.sendMsgToKafka(&msgToMQ, msgToMQ.MsgData.SendID)

		//callbackAfterSendSingle(pb)
	case constant.GroupChatType:

	}
	return reply, nil
}

func callbackAfterSendSingle(p any) {

}

func callbackBeforeSendSingle(req *pb.SendMsgReq) (bool, error) {

	return true, nil
}

func (rpc *RpcChatServer) sendMsgToKafka(m *pb.MsgDataToMQ, key string) error {
	_, _, err := rpc.producer.SendMsg(m, key)
	return err
}
