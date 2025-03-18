package service

import (
	"context"

	"github.com/adnpa/IM/app/user/global"
	"github.com/adnpa/IM/app/user/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/internal/utils"
)

func Model2PB(u model.User) *pb.UserInfo {
	return &pb.UserInfo{
		Id:       u.ID,
		PassWord: u.Passwd,
		Salt:     u.Salt,
		Mobile:   u.Mobile,
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}

type UserService struct {
	pb.UnimplementedUserServer
}

func (s *UserService) GetUserByPage(_ context.Context, in *pb.GetUserByPageReq) (*pb.GetUserByPageResp, error) {
	var users []model.User
	offset := (in.Pn - 1) * in.PSize
	if err := global.DB.Offset(int(offset)).Limit(int(in.PSize)).Find(&users).Error; err != nil {
		return nil, err
	}

	var pbusers []*pb.UserInfo
	for _, u := range users {
		pbusers = append(pbusers, Model2PB(u))
	}
	return &pb.GetUserByPageResp{Usl: pbusers}, nil
}

func (s *UserService) GetUserByMobile(_ context.Context, in *pb.GetUserByMobileReq) (*pb.GetUserByMobileResp, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: in.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return &pb.GetUserByMobileResp{
		Usr: Model2PB(user),
	}, nil
}

func (s *UserService) GetUserByEmail(_ context.Context, in *pb.GetUserByEmailReq) (*pb.GetUserByEmailResp, error) {
	var user model.User
	result := global.DB.Where(&model.User{Email: in.Email}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return &pb.GetUserByEmailResp{
		Usr: Model2PB(user),
	}, nil
}

func (s *UserService) GetUserById(_ context.Context, in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	var user model.User
	result := global.DB.First(&user, in.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.GetUserByIdResp{
		Usr: Model2PB(user),
	}, nil
}

func (s *UserService) GetUserByIds(_ context.Context, in *pb.GetUserByIdsReq) (*pb.GetUserByIdsResp, error) {
	var users []model.User
	result := global.DB.Find(&users, in.Ids)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	var pbUsers []*pb.UserInfo
	for _, user := range users {
		pbUsers = append(pbUsers, Model2PB(user))
	}
	return &pb.GetUserByIdsResp{
		Usrs: pbUsers,
	}, nil
}

func (s *UserService) CreateUser(_ context.Context, in *pb.CreateUserReq) (*pb.CreateUserResp, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: in.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	salt := utils.RandomSalt()
	hashPwd := utils.HashPassword(in.Password, salt)
	user = model.User{
		Nickname: in.Nickname,
		Mobile:   in.Mobile,
		Email:    in.Email,
		Salt:     salt,
		Passwd:   hashPwd,
	}

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	return &pb.CreateUserResp{
		Uid: user.ID,
	}, nil
}

func (s *UserService) UpdateUser(_ context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	var user model.User
	result := global.DB.First(&user, in.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	user.Nickname = in.Nickname
	user.Sex = in.Sex

	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &pb.UpdateUserResp{}, nil
}

func (s *UserService) DeleteUser(_ context.Context, in *pb.DeleteUserReq) (*pb.DeleteUserResp, error) {
	result := global.DB.Delete(&model.User{}, in.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return &pb.DeleteUserResp{}, nil
}

func (s *UserService) CheckPassWord(_ context.Context, in *pb.CheckPassWordReq) (*pb.CheckPassWordResp, error) {
	return &pb.CheckPassWordResp{
		Match: utils.DoPasswordsMatch(in.EncryptedPassword, in.Password, in.Salt),
	}, nil
}
