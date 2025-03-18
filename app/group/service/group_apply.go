package service

import (
	"context"
	"errors"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/group/global"
	"github.com/adnpa/IM/app/group/model"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func GroupApplyModel2PB(g model.GroupApply) *pb.GroupApply {
	return &pb.GroupApply{
		GroupId:     g.GroupID,
		ApplicantId: g.ApplicantID,
		HandlerId:   g.HandlerID,
		Status:      g.Status,
	}
}

// 申请管理
func (s *GroupService) GetGroupApplyById(_ context.Context, in *pb.GetGroupApplyByGroupIdReq) (*pb.GetGroupApplyByGroupIdResp, error) {
	var applies []model.GroupApply
	result := global.DB.Where("group_id = ?", in.GroupId).Find(&applies)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	var pbApplies []*pb.GroupApply
	for _, apply := range applies {
		pbApplies = append(pbApplies, GroupApplyModel2PB(apply))
	}
	return &pb.GetGroupApplyByGroupIdResp{
		GroupApplyList: pbApplies,
	}, nil
}

// func (s *GroupService) GetGroupApplyByIds(_ context.Context, in *pb.GetGroupApplyByIdsReq) (*pb.GetGroupApplyByIdsResp, error) {
// 	panic("not implemented") // TODO: Implement
// }

func (s *GroupService) CreateGroupApply(_ context.Context, in *pb.CreateGroupApplyReq) (*pb.CreateGroupApplyResp, error) {
	var apply model.GroupApply
	copier.Copy(&apply, in)
	result := global.DB.Create(&apply)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &pb.CreateGroupApplyResp{
		Succ: true,
	}, nil
}

func (s *GroupService) UpdateGroupApply(_ context.Context, in *pb.UpdateGroupApplyReq) (*pb.UpdateGroupApplyResp, error) {
	var apply model.GroupApply
	result := global.DB.First(&apply, in.Apply.GroupId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "申请不存在")
		}
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	copier.Copy(&apply, in)

	result = global.DB.Save(&apply)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &pb.UpdateGroupApplyResp{
		Succ: true,
	}, nil
}

func (s *GroupService) DeleteGroupApply(_ context.Context, in *pb.DeleteGroupApplyReq) (*pb.DeleteGroupApplyResp, error) {
	result := global.DB.Delete(&model.GroupApply{}, in.GroupId)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "申请不存在")
	}
	return &pb.DeleteGroupApplyResp{
		Succ: true,
	}, nil
}
