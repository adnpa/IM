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
	// var apply model.FriendApply
	// result := global.DB.First(&apply, req.Id)
	// if result.Error != nil {
	// 	return nil, status.Errorf(codes.NotFound, "申请不存在")
	// }

	// return &pb.GetFriendApplyResp{Apply: &pb.FriendApply{
	// 	Id:          apply.ID,
	// 	FromId:      apply.FromID,
	// 	ToId:        apply.ToID,
	// 	Status:      apply.Status,
	// 	ApplyReason: apply.ApplyReason,
	// 	CreatedAt:   apply.CreatedAt.Format(time.RFC3339),
	// 	UpdatedAt:   apply.UpdatedAt.Format(time.RFC3339),
	// }}, nil
	return &pb.GetFriendApplyResp{}, nil
}

func (s *FriendService) GetFriendApplyByFromId(_ context.Context, req *pb.GetFriendApplyByFromIdReq) (*pb.GetFriendApplyByFromIdResp, error) {
	var applies []model.FriendApply
	result := global.DB.Where(&model.FriendApply{FromID: req.FromId}).Find(&applies)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "查询失败: %v", result.Error)
	}

	var pbApplies []*pb.FriendApply
	for _, a := range applies {
		pbApplies = append(pbApplies, FriendApplyModel2PB(a))
	}

	return &pb.GetFriendApplyByFromIdResp{FriendApply: pbApplies}, nil
}

func (s *FriendService) GetFriendApplyByToId(_ context.Context, req *pb.GetFriendApplyByToIdReq) (*pb.GetFriendApplyByToIdResp, error) {
	var applies []model.FriendApply
	result := global.DB.Where(&model.FriendApply{ToID: req.ToId}).Find(&applies)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "查询失败: %v", result.Error)
	}

	var pbApplies []*pb.FriendApply
	for _, a := range applies {
		pbApplies = append(pbApplies, FriendApplyModel2PB(a))
	}

	return &pb.GetFriendApplyByToIdResp{FriendApply: pbApplies}, nil
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
