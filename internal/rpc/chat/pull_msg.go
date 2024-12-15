package chat

import (
	"context"
	pb2 "github.com/adnpa/IM/pkg/pb"
)

func (rpc *RpcChatServer) GetMaxAndMinSeq(_ context.Context, in *pb2.GetMaxAndMinSeqReq) (*pb2.GetMaxAndMinSeqResp, error) {

}

func (rpc *rpcChat) PullMessageBySeqList(_ context.Context, in *pb2.PullMessageBySeqListReq) (*pb2.PullMessageBySeqListResp, error) {

	//maxSeq, err := redis.GetUserMaxSeq(in.UserID)
}
