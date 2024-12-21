package chat

import (
	"context"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/adnpa/IM/pkg/common/db/redis"
	"github.com/adnpa/IM/pkg/pb/pb_chat"
	"github.com/adnpa/IM/pkg/pb/pb_ws"
)

func (rpc *RpcChatServer) GetMaxAndMinSeq(_ context.Context, in *pb_chat.GetMaxAndMinSeqReq) (*pb_chat.GetMaxAndMinSeqResp, error) {
	uid := in.UserID
	maxSeq, err := redis.GetUserMaxSeq(uid)
	if err != nil {
		return nil, err
	}
	minSeq, err := redis.GetUserMinSeq(uid)
	if err != nil {
		return nil, err
	}

	return &pb_chat.GetMaxAndMinSeqResp{
		ErrCode: 0,
		ErrMsg:  "",
		MaxSeq:  uint32(maxSeq),
		MinSeq:  uint32(minSeq),
	}, nil
}

// PullMessageBySeqList 离线消息拉取/超时重传
func (rpc *RpcChatServer) PullMessageBySeqList(_ context.Context, in *pb_ws.PullMessageBySeqListReq) (*pb_ws.PullMessageBySeqListResp, error) {
	msgList, err := mongodb.GetMsgBySeqList(in.UserID, in.SeqList)
	if err != nil {
		return &pb_ws.PullMessageBySeqListResp{
			ErrCode: constant.ErrChatMsgTimeout,
			ErrMsg:  constant.StatusText(constant.ErrChatMsgTimeout),
		}, err
	}
	//maxSeq, err := redis.GetUserMaxSeq(in.UserID)
	return &pb_ws.PullMessageBySeqListResp{
		ErrCode: 0,
		ErrMsg:  "",
		List:    msgList,
	}, nil
}
