package service

import (
	"context"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/friend/global"
	"github.com/adnpa/IM/app/friend/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FriendApplyModel2PB(apply model.FriendApply) *pb.FriendApply {
	return &pb.FriendApply{
		From: apply.FromID,
		To:   apply.ToID,
		Flag: apply.Status,
		// ApplyReason: apply.ApplyReason,
	}
}

func (s *FriendService) GetFriendApply(_ context.Context, req *pb.GetFriendApplyReq) (*pb.GetFriendApplyResp, error) {
	var friendApplys []model.FriendApply
	var pbFriendApplys []*pb.FriendApply
	err := global.DB.Where(&model.FriendApply{FromID: req.UserId}).Or(&model.FriendApply{ToID: req.UserId}).Find(&friendApplys).Error
	if err != nil {
		return nil, err
	}

	for _, i := range friendApplys {
		pbFriendApplys = append(pbFriendApplys, FriendApplyModel2PB(i))
	}
	return &pb.GetFriendApplyResp{
		FriendApplys: pbFriendApplys,
	}, nil
}

func (s *FriendService) CreateFriendApply(_ context.Context, req *pb.CreateFriendApplyReq) (*pb.CreateFriendApplyResp, error) {
	apply := model.FriendApply{
		FromID: req.FriendApply.From,
		ToID:   req.FriendApply.To,
	}

	result := global.DB.Create(&apply)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "创建失败: %v", result.Error)
	}

	return &pb.CreateFriendApplyResp{Success: true}, nil
}

func (s *FriendService) UpdateFriendApply(_ context.Context, req *pb.UpdateFriendApplyReq) (*pb.UpdateFriendApplyResp, error) {
	// result := global.DB.Model(&model.FriendApply{}).Where("id = ?", req.FriendApply.From).Update("status", req.Status)
	// if result.Error != nil {
	// 	return nil, status.Errorf(codes.Internal, "更新失败: %v", result.Error)
	// }

	return &pb.UpdateFriendApplyResp{Success: true}, nil
}

func (s *FriendService) DeleteFriendApply(_ context.Context, req *pb.DeleteFriendApplyReq) (*pb.DeleteFriendApplyResp, error) {
	result := global.DB.Delete(&model.FriendApply{}, req.FromId)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "删除失败: %v", result.Error)
	}

	return &pb.DeleteFriendApplyResp{Success: true}, nil
}
