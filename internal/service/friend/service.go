package friend

import (
	"context"

	"github.com/adnpa/IM/internal/model"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/adnpa/IM/pkg/common/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

// friend
// apply
// s *FriendService pb.FriendServer

func Model2PB(model.Friend) *pb.FriendInfo {
	return &pb.FriendInfo{}
}

type FriendService struct {
	pb.UnimplementedFriendServer
}

type AddFriendReq struct {
	OwnId int64 `json:"own_id"`
	FriId int64 `json:"fri_id"`
}

func (s *FriendService) GetFriendsById(_ context.Context, in *pb.GetFriendsByIdReq) (*pb.GetFriendsByIdResp, error) {
	var friendL []*pb.FriendInfo
	cur, err := mongodb.GetAll("friend", bson.M{"ownerid": in.Ownid})
	if err != nil {
		logger.Error("", zap.Error(err))
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var tmp model.Friend
		cur.Decode(&tmp)
		friendL = append(friendL, Model2PB(tmp))
	}

	return &pb.GetFriendsByIdResp{
		Friends: friendL,
	}, nil
}

func (s *FriendService) CreateFriend(_ context.Context, in *pb.CreateFriendReq) (*pb.CreateFriendResp, error) {
	return &pb.CreateFriendResp{}, mongodb.Insert("friend", &model.Friend{OwnerID: in.Info.Owner, FriendID: in.Info.Friend})
}

func (s *FriendService) DeleteFriend(_ context.Context, _ *pb.DeleteFriendReq) (*pb.DeleteFriendResp, error) {
	// mongodb.Delete("friend", "")
	panic("not implemented") // TODO: Implement
}

func (s *FriendService) UpdateFriend(_ context.Context, _ *pb.UpdateFriendReq) (*pb.UpdateFriendResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *FriendService) GetFriendApplyById(_ context.Context, _ *pb.GetFriendApplyByIdReq) (*pb.GetFriendApplyByIdResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *FriendService) GetFriendApplyByIds(_ context.Context, _ *pb.GetFriendApplyByIdsReq) (*pb.GetFriendApplyByIdsResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *FriendService) CreateFriendApply(_ context.Context, _ *pb.CreateFriendApplyReq) (*pb.CreateFriendApplyResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *FriendService) UpdateFriendApply(_ context.Context, _ *pb.UpdateFriendApplyReq) (*pb.UpdateFriendApplyResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *FriendService) DeleteFriendApply(_ context.Context, _ *pb.DeleteFriendApplyReq) (*pb.DeleteFriendApplyResp, error) {
	panic("not implemented") // TODO: Implement
}
