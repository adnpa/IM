package service

import (
	"context"
	"errors"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/group/constant"
	"github.com/adnpa/IM/app/group/global"
	"github.com/adnpa/IM/app/group/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// s *GroupService pb.GroupServer

func GroupModel2PB(g model.Group) *pb.GroupInfo {
	return &pb.GroupInfo{
		GroupId:   g.GroupID,
		GroupName: g.GroupName,
		CreatorId: g.CreatorID,
	}
}

type GroupService struct {
	pb.UnimplementedGroupServer
}

func (s *GroupService) GetUserGroups(_ context.Context, in *pb.GetUserGroupsReq) (*pb.GetUserGroupsResp, error) {
	var groups []model.GroupMember
	var gids []int64
	err := global.DB.Where(&model.GroupMember{UserID: in.UserId}).Find(&groups).Error
	if err != nil {
		return nil, err
	}

	for _, g := range groups {
		gids = append(gids, g.GroupID)
	}
	return &pb.GetUserGroupsResp{GroupIds: gids}, nil
}

// 群聊基础信息管理
func (s *GroupService) GetGroupInfoById(_ context.Context, in *pb.GetGroupInfoByIdReq) (*pb.GetGroupInfoByIdResp, error) {
	var group model.Group
	result := global.DB.First(&group, in.GroupId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "群聊不存在")
		}
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &pb.GetGroupInfoByIdResp{
		GroupInfo: GroupModel2PB(group),
	}, nil
}

func (s *GroupService) CreateGroupInfo(_ context.Context, in *pb.CreateGroupInfoReq) (*pb.CreateGroupInfoResp, error) {
	g := model.Group{
		GroupID:     in.GroupInfo.GroupId,
		GroupName:   in.GroupInfo.GroupName,
		CreatorID:   in.GroupInfo.CreatorId,
		AvatarURL:   in.GroupInfo.AvatarUrl,
		Description: in.GroupInfo.Description,
		MaxMembers:  in.GroupInfo.MaxMembers,
	}

	gm := model.GroupMember{
		GroupID: in.GroupInfo.GroupId,
		UserID:  in.GroupInfo.CreatorId,
		Role:    constant.RoleOwner,
	}

	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Create(&g).Error; err != nil {
			return err
		}

		gm.GroupID = g.GroupID
		if err := tx.Create(&gm).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "create group fail:%v", err)
	}
	return &pb.CreateGroupInfoResp{
		GroupId: g.GroupID,
	}, nil
}

func (s *GroupService) UpdateGroupInfo(_ context.Context, in *pb.UpdateGroupInfoReq) (*pb.UpdateGroupInfoResp, error) {
	var group model.Group
	result := global.DB.First(&group, in.GroupInfo.GroupId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "群聊不存在")
		}
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	// update
	group.Description = in.GroupInfo.Description

	result = global.DB.Save(&group)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &pb.UpdateGroupInfoResp{
		Succ: true,
	}, nil
}

func (s *GroupService) DeleteGroupInfo(_ context.Context, in *pb.DeleteGroupInfoReq) (*pb.DeleteGroupInfoResp, error) {
	result := global.DB.Delete(&model.Group{}, in.GroupId)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "群聊不存在")
	}
	return &pb.DeleteGroupInfoResp{
		Succ: true,
	}, nil
}
