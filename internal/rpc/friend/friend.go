package friend

import (
	"context"
	"errors"
	utils2 "github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/model"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/adnpa/IM/pkg/common/db/mysql/dao"
	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/adnpa/IM/pkg/discovery"
	"github.com/adnpa/IM/pkg/pb/pb_friend"
	"github.com/adnpa/IM/pkg/pb/pb_ws"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"strings"
)

type RpcFriendServer struct {
	*pb_friend.UnimplementedFriendServer

	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func NewFriendServer(port int) *RpcFriendServer {
	return &RpcFriendServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.FriendName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
}

func (rpc *RpcFriendServer) Run() {
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
	pb_friend.RegisterFriendServer(srv, rpc)
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

func (rpc *RpcFriendServer) AddFriend(ctx context.Context, req *pb_friend.AddFriendReq) (*pb_ws.CommonResp, error) {
	if _, err := dao.GetUserByUid(req.Uid); err != nil {
		return nil, err
	}

	friendReq := &model.FriendRequest{
		ReqID:      req.Uid,
		UserID:     req.FriendUid,
		ReqMessage: req.ReqMessage,
	}
	err := dao.AddFriendRequest(friendReq)
	if err != nil {
		return nil, err
	}

	//todo 推送被申请者
	return &pb_ws.CommonResp{ErrorMsg: "", ErrorCode: 0}, nil
}

func (rpc *RpcFriendServer) GetFriendApplyList(ctx context.Context, req *pb_friend.GetFriendApplyReq) (*pb_friend.GetFriendApplyResp, error) {
	list, err := dao.GetApplyList(req.Uid)
	if err != nil {
		logger.L().Info("GetFriendApplyList db err", zap.Error(err))
		return nil, err
	}

	var data []*pb_friend.ApplyUserInfo

	for _, i := range list {
		tmp := &pb_friend.ApplyUserInfo{
			Name: i.ReqID,
		}
		data = append(data, tmp)
	}

	return &pb_friend.GetFriendApplyResp{
		Data: data,
	}, nil
}

func (rpc *RpcFriendServer) GetSelfApplyList(ctx context.Context, req *pb_friend.GetFriendApplyReq) (*pb_friend.GetFriendApplyResp, error) {
	uid := req.Uid
	list, err := dao.GetSelfApplyList(uid)
	if err != nil {
		logger.L().Info("GetSelfApplyList db err", zap.Error(err))
		return nil, err
	}

	var data []*pb_friend.ApplyUserInfo

	for _, i := range list {
		tmp := &pb_friend.ApplyUserInfo{
			Name: i.ReqID,
		}
		data = append(data, tmp)
	}

	return &pb_friend.GetFriendApplyResp{
		Data: data,
	}, nil
}

// AddFriendResponse 添加好友回复
func (rpc *RpcFriendServer) AddFriendResponse(ctx context.Context, req *pb_friend.AddFriendResponseReq) (*pb_ws.CommonResp, error) {
	fq, err := dao.GetFriendReq(req.FriendUid, req.Uid)
	if err != nil {
		return nil, err
	}
	fq.Flag = req.Flag
	//todo 推送好友

	if req.Flag == constant.FriendAgreeFlag {
		err = dao.AddFriend(fq)
	} else if req.Flag == constant.FriendRefuseFlag {
		err = dao.UpdateFriendRequest(fq)
	} else {
		err = errors.New("error flag")
	}

	return &pb_ws.CommonResp{}, nil
}

func (rpc *RpcFriendServer) IsFriend(ctx context.Context, req *pb_friend.IsFriendReq) (*pb_friend.IsFriendResp, error) {
	uid := req.Uid

	isFriend := dao.IsFriend(uid, req.FriendUid)

	var flag int32
	if isFriend {
		flag = constant.FriendAgreeFlag
	} else {
		flag = constant.FriendRefuseFlag
	}
	return &pb_friend.IsFriendResp{ShipType: flag}, nil
}

func (rpc *RpcFriendServer) GetFriendList(ctx context.Context, req *pb_friend.GetFriendListReq) (*pb_friend.GetFriendListResp, error) {
	var data []*pb_ws.UserInfo
	friends, err := dao.GetFriendsByUserUid(req.Uid)
	if err != nil {
		return nil, err
	}
	for _, f := range friends {
		friend := &pb_ws.UserInfo{
			UserID: f.OwnerID,
			//f.
			//Uid:     f.FriendUID,
			//Comment: f.Comment,
		}
		data = append(data, friend)
	}
	return &pb_friend.GetFriendListResp{
		Data: data,
	}, nil
}

func (rpc *RpcFriendServer) SetFriendComment(ctx context.Context, req *pb_friend.SetFriendCommentReq) (*pb_ws.CommonResp, error) {
	err := dao.SetComment(req.Uid, req.FriendUid, req.Comment)
	return &pb_ws.CommonResp{}, err
}

func (rpc *RpcFriendServer) DeleteFriend(ctx context.Context, req *pb_friend.DeleteFriendReq) (*pb_ws.CommonResp, error) {
	err := dao.DeleteFriend(req.Uid, req.FriendUid)
	if err != nil {
		return nil, err
	}
	return &pb_ws.CommonResp{}, nil
}

func (rpc *RpcFriendServer) AddBlacklist(ctx context.Context, req *pb_friend.AddBlacklistReq) (*pb_ws.CommonResp, error) {
	err := dao.AddBlacklist(req.Uid, req.FriendUid)
	return &pb_ws.CommonResp{}, err
}

func (rpc *RpcFriendServer) RemoveBlacklist(ctx context.Context, req *pb_friend.RemoveBlacklistReq) (*pb_ws.CommonResp, error) {
	err := dao.RemoveBlacklist(req.Uid, req.FriendUid)
	return &pb_ws.CommonResp{}, err
}

func (rpc *RpcFriendServer) IsInBlackList(ctx context.Context, req *pb_friend.IsInBlackListReq) (*pb_friend.IsInBlackListResp, error) {
	boo := dao.IsInBlackList(req.SendUid, req.ReceiveUid)
	return &pb_friend.IsInBlackListResp{Response: boo}, nil
}
func (rpc *RpcFriendServer) GetBlacklist(ctx context.Context, req *pb_friend.GetBlacklistReq) (*pb_friend.GetBlacklistResp, error) {
	blacklist, err := dao.GetBlacklist(req.Uid)
	if err != nil {
		return &pb_friend.GetBlacklistResp{}, err
	}
	var data []*pb_ws.UserInfo
	for _, i := range blacklist {
		tmp := &pb_ws.UserInfo{
			UserID: i.OwnerID,
		}
		data = append(data, tmp)
	}
	return &pb_friend.GetBlacklistResp{Data: data}, nil
}

// ImportFriend 管理员批量导入好友
func (rpc *RpcFriendServer) ImportFriend(ctx context.Context, req *pb_friend.ImportFriendReq) (*pb_friend.ImportFriendResp, error) {
	return nil, nil
}
