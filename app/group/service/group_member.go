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

func GroupMemberModel2PB(u model.GroupMember) *pb.GroupMember {
	return &pb.GroupMember{
		Id:     u.ID,
		UserId: u.UserID,
		Role:   u.Role,
	}
}

// 成员管理
func (s *GroupService) GetGroupMemberById(_ context.Context, in *pb.GetGroupMemberByIdReq) (*pb.GetGroupMemberByIdResp, error) {
	var members []model.GroupMember
	result := global.DB.Where("group_id = ?", in.GroupId).Find(&members)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	var pbMembers []*pb.GroupMember
	for _, member := range members {
		pbMembers = append(pbMembers, GroupMemberModel2PB(member))
	}
	return &pb.GetGroupMemberByIdResp{
		Members: pbMembers,
	}, nil
}

func (s *GroupService) CreateGroupMember(_ context.Context, in *pb.CreateGroupMemberReq) (*pb.CreateGroupMemberResp, error) {
	var member model.GroupMember
	copier.Copy(&member, in)

	result := global.DB.Create(&member)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &pb.CreateGroupMemberResp{
		Succ: true,
	}, nil
}

func (s *GroupService) UpdateGroupMember(_ context.Context, in *pb.UpdateGroupMemberReq) (*pb.UpdateGroupMemberResp, error) {
	var member model.GroupMember
	result := global.DB.First(&member, in.Member.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "成员不存在")
		}
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	copier.Copy(&member, in)

	result = global.DB.Save(&member)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &pb.UpdateGroupMemberResp{
		Succ: true,
	}, nil
}

func (s *GroupService) DeleteGroupMember(_ context.Context, in *pb.DeleteGroupMemberReq) (*pb.DeleteGroupMemberResp, error) {
	result := global.DB.Delete(&model.GroupMember{}, in.GroupId)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "成员不存在")
	}
	return &pb.DeleteGroupMemberResp{
		Succ: true,
	}, nil
}
