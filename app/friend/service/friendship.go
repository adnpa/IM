package service

import (
	"context"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/friend/global"
	"github.com/adnpa/IM/app/friend/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// friend
// apply
// s *FriendService pb.FriendServer

func FriendshipModel2PB(friendship model.Friendship) *pb.Friendship {
	return &pb.Friendship{
		UserId:   friendship.UserID,
		FriendId: friendship.FriendID,
		Comment:  friendship.Comment,
	}
}

type FriendService struct {
	pb.UnimplementedFriendServer
}

func (s *FriendService) GetFriendsByUserId(_ context.Context, _ *pb.GetFriendsByUserIdReq) (*pb.GetFriendsByUserIdResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *FriendService) CreateFriend(_ context.Context, in *pb.CreateFriendReq) (*pb.CreateFriendResp, error) {
	friendship := model.Friendship{
		UserID:   in.Info.UserId,
		FriendID: in.Info.FriendId,
		Comment:  in.Info.Comment,
	}

	result := global.DB.Create(&friendship)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "创建失败: %v", result.Error)
	}

	return &pb.CreateFriendResp{Success: true}, nil
}

func (s *FriendService) DeleteFriend(_ context.Context, req *pb.DeleteFriendReq) (*pb.DeleteFriendResp, error) {
	result := global.DB.Where(&model.Friendship{UserID: req.Uid, FriendID: req.FriendId}).Delete(&model.Friendship{})
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "删除失败: %v", result.Error)
	}

	return &pb.DeleteFriendResp{Success: true}, nil
}

func (s *FriendService) UpdateFriend(_ context.Context, req *pb.UpdateFriendReq) (*pb.UpdateFriendResp, error) {
	result := global.DB.Model(&model.Friendship{}).Where(&model.Friendship{UserID: req.Info.UserId, FriendID: req.Info.FriendId}).Update("comment", req.Info.Comment)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "更新失败: %v", result.Error)
	}

	return &pb.UpdateFriendResp{}, nil
}

func (s *FriendService) mustEmbedUnimplementedFriendServer() {
	panic("not implemented") // TODO: Implement
}
