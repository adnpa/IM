package service

import (
	"context"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/offline/model"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
)

// s *OfflineService pb.OfflineServer

var COLL_NAME = "offline"

type OfflineService struct {
	pb.UnimplementedOfflineServer
}

func (s *OfflineService) GetOfflineMsg(_ context.Context, in *pb.GetOfflineMsgReq) (*pb.GetOfflineMsgResp, error) {

	return &pb.GetOfflineMsgResp{}, nil
}

func (s *OfflineService) PutMsg(_ context.Context, in *pb.PutMsgReq) (*pb.PutMsgResp, error) {
	var msg model.Message
	copier.Copy(&msg, in)
	// todo inbox logic
	// userId := in.UserId
	err := mongodb.Insert(COLL_NAME, msg)
	if err != nil {
		return nil, err
	}
	return &pb.PutMsgResp{Succ: true}, nil
}

func (s *OfflineService) RemoveMsg(_ context.Context, in *pb.RemoveMsgReq) (*pb.RemoveMsgResp, error) {
	for _, id := range in.MsgIds {
		err := mongodb.Delete(COLL_NAME, bson.M{"id": id})
		if err != nil {
			return nil, err
		}
	}
	return &pb.RemoveMsgResp{Succ: true}, nil
}
