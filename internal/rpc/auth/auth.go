package auth

import (
	"context"
	"errors"
	utils2 "github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/db/mysql/dao"
	"github.com/adnpa/IM/pkg/common/db/mysql/model"
	"github.com/adnpa/IM/pkg/discovery"
	"github.com/adnpa/IM/pkg/pb/pb_auth"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"strings"
)

type RpcAuthServer struct {
	*pb_auth.UnimplementedAuthServer

	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func NewAuthServer(port int) *RpcAuthServer {
	return &RpcAuthServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.AuthName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
}

func (rpc *RpcAuthServer) Run() {
	address := utils2.ServerIP + ":" + strconv.Itoa(rpc.rpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	//grpc server
	srv := grpc.NewServer()
	defer srv.GracefulStop()

	//service registers with etcd
	pb_auth.RegisterAuthServer(srv, rpc)
	err = discovery.Register(rpc.etcdSchema, strings.Join(rpc.etcdAddr, ","), utils2.ServerIP, rpc.rpcPort, rpc.rpcRegisterName, 10)
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

func (rpc *RpcAuthServer) Register(ctx context.Context, params *pb_auth.RegisterReq) (*pb_auth.RegisterResp, error) {
	user := &model.User{
		UID:    params.UID,
		Name:   params.Name,
		Icon:   params.Icon,
		Gender: params.Gender,
		Mobile: params.Mobile,
		Birth:  params.Birth,
		Email:  params.Email,
	}
	err := dao.CreateUser(user)
	if err != nil {
		return &pb_auth.RegisterResp{Success: false}, err
	}
	return &pb_auth.RegisterResp{Success: true}, nil
}

func (rpc *RpcAuthServer) Token(ctx context.Context, params *pb_auth.TokenReq) (*pb_auth.TokenResp, error) {
	_, err := dao.GetUserByUid(params.UID)
	if err != nil {
		return &pb_auth.TokenResp{ErrCode: 500, ErrMsg: "not user"}, errors.New("")
	}

	token, ex, err := utils2.GenerateToken(params.UID, params.Platform)
	if err != nil {
		return &pb_auth.TokenResp{ErrCode: 500, ErrMsg: "not user"}, errors.New("")
	}

	return &pb_auth.TokenResp{
		Token:       token,
		ExpiredTime: ex,
	}, nil
}
