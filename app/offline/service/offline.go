package service

import (
	"context"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/online/model"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/jinzhu/copier"
)

// s *OfflineService pb.OfflineServer

var COLL_NAME = "offline"

type OfflineService struct {
	pb.UnimplementedOfflineServer
}

func (s *OfflineService) GetOfflineMsg(_ context.Context, _ *pb.GetOfflineMsgReq) (*pb.GetOfflineMsgResp, error) {
	
}

func (s *OfflineService) PutMsg(_ context.Context, in *pb.PutMsgReq) (*pb.PutMsgResp, error) {
	var msg model.Message
	copier.Copy(&msg, in)

	err := mongodb.Insert(COLL_NAME, msg)
	if err != nil {
		return nil, err
	}
	return &pb.PutMsgResp{Succ: true}, nil
}

func (s *OfflineService) RemoveMsg(_ context.Context, in *pb.RemoveMsgReq) (*pb.RemoveMsgResp, error) {
	err := mongodb.Delete(COLL_NAME, in.MsgIds)
	if err != nil {
		return nil, err
	}
	return &pb.RemoveMsgResp{Succ: true}, nil
}
