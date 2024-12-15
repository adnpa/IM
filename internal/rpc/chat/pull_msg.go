package chat

import (
	"context"
	"github.com/adnpa/IM/pkg/pb/pb_chat"
	"github.com/adnpa/IM/pkg/pb/pb_ws"
)

func (rpc *RpcChatServer) GetMaxAndMinSeq(_ context.Context, in *pb_chat.GetMaxAndMinSeqReq) (*pb_chat.GetMaxAndMinSeqResp, error) {

}

func (rpc *rpcChat) PullMessageBySeqList(_ context.Context, in *pb_ws.PullMessageBySeqListReq) (*pb_ws.PullMessageBySeqListResp, error) {

	//maxSeq, err := redis.GetUserMaxSeq(in.UserID)
}
