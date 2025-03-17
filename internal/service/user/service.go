package user

import (
	"context"

	"github.com/adnpa/IM/internal/model"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/adnpa/IM/pkg/common/pb"
	"go.mongodb.org/mongo-driver/bson"
)

func Model2PB(u model.User) *pb.UserInfo {
	return &pb.UserInfo{
		Id:       int32(u.Id),
		PassWord: u.Passwd,
		Nickname: u.Nickname,
		Mobile:   u.Mobile,
		Email:    u.Email,
	}
}

type UserService struct {
	pb.UnimplementedUserServer
}

func (s *UserService) GetUserByPage(_ context.Context, _ *pb.GetUserByPageReq) (*pb.GetUserByPageResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *UserService) GetUserByMobile(_ context.Context, in *pb.GetUserByMobileReq) (*pb.GetUserByMobileResp, error) {
	user := model.User{}
	err := mongodb.GetDecode("user", bson.M{"mobile": in.Mobile}, &user)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByMobileResp{
		Usr: Model2PB(user),
	}, nil
}

func (s *UserService) GetUserByEmail(_ context.Context, in *pb.GetUserByEmailReq) (*pb.GetUserByEmailResp, error) {
	u := model.User{}
	mongodb.GetDecode("user", bson.M{"email": in.Email}, &u)
	return &pb.GetUserByEmailResp{Usr: Model2PB(u)}, nil
}

func (s *UserService) GetUserById(_ context.Context, in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	u := model.User{}
	mongodb.GetDecode("user", bson.M{"id": in.Id}, &u)
	return &pb.GetUserByIdResp{Usr: Model2PB(u)}, nil
}

func (s *UserService) GetUserByIds(_ context.Context, _ *pb.GetUserByIdsReq) (*pb.GetUserByIdsResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *UserService) CreateUser(_ context.Context, in *pb.CreateUserReq) (*pb.CreateUserResp, error) {
	salt := utils.RandomSalt()
	hashPwd := utils.HashPassword(in.Password, salt)
	u := model.User{
		Nickname: in.Nickname,
		Mobile:   in.Mobile,
		Email:    in.Email,
		Salt:     salt,
		Passwd:   hashPwd,
	}
	return &pb.CreateUserResp{}, mongodb.Insert("user", &u)
}

func (s *UserService) UpdateUser(_ context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *UserService) DeleteUser(_ context.Context, in *pb.DeleteUserReq) (*pb.DeleteUserResp, error) {
	mongodb.Delete("user", bson.M{"id": in.Id})
	return &pb.DeleteUserResp{}, nil
}

func (s *UserService) CheckPassWord(_ context.Context, in *pb.CheckPassWordReq) (*pb.CheckPassWordResp, error) {
	return &pb.CheckPassWordResp{
		Match: utils.DoPasswordsMatch(in.EncryptedPassword, in.Password, in.Salt),
	}, nil
}

func (s *UserService) mustEmbedUnimplementedUserServer() {
	panic("not implemented") // TODO: Implement
}
