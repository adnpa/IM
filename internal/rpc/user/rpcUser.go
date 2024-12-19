package user

import (
	"context"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/model"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/adnpa/IM/pkg/common/db/mysql/dao"
	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/adnpa/IM/pkg/discovery"
	"github.com/adnpa/IM/pkg/pb/pb_user"
	"github.com/adnpa/IM/pkg/pb/pb_ws"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strconv"
	"strings"
)

type RpcUserServer struct {
	*pb_user.UnimplementedUserServer

	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func NewUserServer(port int) *RpcUserServer {
	return &RpcUserServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.UserName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
}

func (rpc *RpcUserServer) Run() {
	address := utils.ServerIP + ":" + strconv.Itoa(rpc.rpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	//grpc server
	srv := grpc.NewServer()
	defer srv.GracefulStop()

	//service registers with etcd
	pb_user.RegisterUserServer(srv, rpc)
	err = discovery.Register(rpc.etcdSchema, strings.Join(rpc.etcdAddr, ","), utils.ServerIP, rpc.rpcPort, rpc.rpcRegisterName, 10)
	if err != nil {
		log.Println(err)
		return
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("", "", "rpc get_token init success")
}

func (rpc *RpcUserServer) GetUserInfo(ctx context.Context, req *pb_user.GetUserInfoReq) (*pb_user.GetUserInfoResp, error) {
	logger.L().Info("rpc GetUserInfo start", zap.Any("uidList", req.UserIDList))
	var pbUserL []*pb_ws.UserInfo
	if len(req.UserIDList) > 0 {
		userL, err := dao.GetUsersByUidL(req.UserIDList)
		if err != nil {
			return nil, err
		}
		for _, u := range userL {
			pbu := &pb_ws.UserInfo{
				UserID:      u.UID,
				Nickname:    u.Name,
				FaceURL:     u.Icon,
				Gender:      u.Gender,
				PhoneNumber: u.Mobile,
				Birth:       u.Birth,
				Email:       u.Email,
			}
			pbUserL = append(pbUserL, pbu)
		}
	} else {
		return &pb_user.GetUserInfoResp{
			ErrorCode: constant.ErrUserArgs,
			ErrorMsg:  "UserIDList is empty",
		}, nil
	}
	return &pb_user.GetUserInfoResp{
		Data: pbUserL,
	}, nil
}

func (rpc *RpcUserServer) UpdateUserInfo(ctx context.Context, req *pb_user.UpdateUserInfoReq) (*pb_ws.CommonResp, error) {
	//todo 权限分离
	//用户自身修改自己
	//appManager修改任意用户
	logger.L().Info("rpc UpdateUserInfo", zap.String("new user info", req.String()))
	err := dao.UpdateUser(&model.User{
		UID:    req.Uid,
		Name:   req.Name,
		Icon:   req.Icon,
		Gender: req.Gender,
		Mobile: req.Mobile,
		Birth:  req.Birth,
		Email:  req.Email,
		//UpdatedAt: time.Now(),
	})
	if err != nil {
		return &pb_ws.CommonResp{ErrorCode: 999, ErrorMsg: "update_err"}, err
	}

	//todo friend

	return &pb_ws.CommonResp{ErrorCode: 0, ErrorMsg: ""}, err
}

func (rpc *RpcUserServer) DeleteUsers(ctx context.Context, req *pb_user.DeleteUsersReq) (*pb_user.DeleteUsersResp, error) {
	// todo appManager批量删除

	return nil, status.Errorf(codes.Unimplemented, "method DeleteUsers not implemented")
}
