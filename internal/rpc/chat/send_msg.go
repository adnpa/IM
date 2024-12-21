package chat

import (
	"context"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/adnpa/IM/pkg/pb/pb_chat"
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

// option

func (rpc *RpcChatServer) SendMsg(_ context.Context, req *pb_chat.SendMsgReq) (*pb_chat.SendMsgResp, error) {

	//isHistory := utils.GetSwitchFromOptions(pb.MsgData.Options, constant.IsHistory)

	switch req.MsgData.SessionType {
	case constant.SingleChatType:
		//canSend, err := callbackBeforeSendSingle(req)
		//if err != nil {
		//	return nil, err
		//}
		// todo 聊天选项控制
		//isSend := modifyMessageByUserMessageReceiveOpt(req.MsgData.RecvID, req.MsgData.SendID, constant.SingleChatType, req)

		msgToMQ := pb_chat.MsgDataToMQ{Token: req.Token, OperationID: req.OperationID, MsgData: req.MsgData}
		err := rpc.sendMsgToKafka(&msgToMQ, msgToMQ.MsgData.SendID)
		if err != nil {
			return &pb_chat.SendMsgResp{
				ErrCode:     constant.ErrChatKafkaSend,
				ErrMsg:      constant.StatusText(constant.ErrChatKafkaSend),
				ServerMsgID: msgToMQ.MsgData.ServerMsgID,
				ClientMsgID: req.MsgData.ClientMsgID,
				SendTime:    msgToMQ.MsgData.SendTime,
			}, err
		}

		return &pb_chat.SendMsgResp{
			ErrCode:     0,
			ErrMsg:      "",
			ServerMsgID: "",
			ClientMsgID: "",
			SendTime:    0,
		}, err
		//callbackAfterSendSingle(pb)
	case constant.GroupChatType:
		return nil, nil
	default:
		reply := &pb_chat.SendMsgResp{
			ErrCode:     constant.ErrChatUnknownMsgType,
			ErrMsg:      constant.StatusText(constant.ErrChatUnknownMsgType),
			ServerMsgID: "",
			ClientMsgID: "",
			SendTime:    0,
		}
		return reply, nil
	}
}

//func modifyMessageByUserMessageReceiveOpt(userID, sourceID string, sessionType int, pb *pb_chat.SendMsgReq) bool {
//	//conversationId := getConversationIDBySessionType(sourceID, sessionType)
//	//redis.
//
//}

//func getConversationIDBySessionType(sourceID string, sessionType int) string {
//	switch sessionType {
//	case constant.SingleChatType:
//		return "single_" + sourceID
//	case constant.GroupChatType:
//		return "group_" + sourceID
//	}
//	return ""
//}

//func callbackAfterSendSingle(p any) {
//
//}

//func callbackBeforeSendSingle(req *pb_chat.SendMsgReq) (bool, error) {
//
//	return true, nil
//}

func (rpc *RpcChatServer) sendMsgToKafka(m *pb_chat.MsgDataToMQ, key string) error {
	_, _, err := rpc.producer.SendMsg(m, key)
	return err
}
