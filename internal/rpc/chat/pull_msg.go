package chat

import (
	"context"
	"github.com/adnpa/IM/pkg/pb/pb_chat"
	"github.com/adnpa/IM/pkg/pb/pb_ws"
)

func (rpc *RpcChatServer) GetMaxAndMinSeq(_ context.Context, in *pb_chat.GetMaxAndMinSeqReq) (*pb_chat.GetMaxAndMinSeqResp, error) {

	return &pb_chat.GetMaxAndMinSeqResp{
		ErrCode: 0,
		ErrMsg:  "",
		MaxSeq:  0,
		MinSeq:  0,
	}, nil
}

func (rpc *RpcChatServer) PullMessageBySeqList(_ context.Context, in *pb_ws.PullMessageBySeqListReq) (*pb_ws.PullMessageBySeqListResp, error) {

	//maxSeq, err := redis.GetUserMaxSeq(in.UserID)
	return &pb_ws.PullMessageBySeqListResp{
		ErrCode: 0,
		ErrMsg:  "",
		List:    nil,
	}, nil
}
