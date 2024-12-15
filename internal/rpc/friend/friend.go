package friend

import (
	"context"
	utils2 "github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/db/mysql/dao"
	"github.com/adnpa/IM/pkg/common/db/mysql/model"
	"github.com/adnpa/IM/pkg/discovery"
	"github.com/adnpa/IM/pkg/pb/pb_friend"
	"github.com/adnpa/IM/pkg/pb/pb_ws"
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

//func (rpc *RpcFriendServer) GetFriendsInfo(context.Context, *pb_friend.GetFriendsInfoReq) (*pb_friend.GetFriendInfoResp, error) {
//	//
//	return nil, nil
//}

func (rpc *RpcFriendServer) GetFriendList(ctx context.Context, req *pb_friend.GetFriendListReq) (*pb_friend.GetFriendListResp, error) {
	//var data []*pb_friend.UserInfo
	//claims, _ := utils2.ParseToken(req.Token)
	//friends, err := dao.GetFriendsByUserUid(claims.UID)
	//if err != nil {
	//	return nil, err
	//}
	//for _, f := range friends {
	//	friend := &pb_friend.UserInfo{
	//		//Uid:     f.FriendUID,
	//		//Comment: f.Comment,
	//	}
	//	data = append(data, friend)
	//}
	return &pb_friend.GetFriendListResp{
		//Data: data,
	}, nil
}

func (rpc *RpcFriendServer) AddFriend(ctx context.Context, req *pb_friend.AddFriendReq) (*pb_ws.CommonResp, error) {

	if _, err := dao.GetUserByUid(req.Uid); err != nil {
		return nil, err
	}
	uid, _ := utils2.GetUserId(req.Token)

	friendReq := &model.FriendRequest{
		ReqID:      uid,
		UserID:     req.Uid,
		ReqMessage: req.ReqMessage,
	}
	err := dao.AddFriendRequest(friendReq)
	if err != nil {
		return nil, err
	}

	//todo 推送被申请者
	return &pb_ws.CommonResp{ErrorMsg: "", ErrorCode: 0}, nil
}

// AddFriendResponse 添加好友回复
func (rpc *RpcFriendServer) AddFriendResponse(ctx context.Context, req *pb_friend.AddFriendResponseReq) (*pb_ws.CommonResp, error) {
	//userid, _ := utils2.GetUserId(req.Token)
	//fq, err := dao.GetFriendReq(req.Uid, userid)
	//if err != nil {
	//	return nil, err
	//}
	//fq.Flag = req.Flag
	//todo 推送好友

	//if req.Flag ==

	return nil, nil
}

func (rpc *RpcFriendServer) DeleteFriend(context.Context, *pb_friend.DeleteFriendReq) (*pb_ws.CommonResp, error) {
	return nil, nil
}

//func (rpc *RpcFriendServer) GetFriendApplyList(context.Context, *pb_friend.GetFriendApplyReq) (*pb_friend.GetFriendApplyResp, error) {
//	return nil, nil
//}
//func (rpc *RpcFriendServer) GetSelfApplyList(context.Context, *pb_friend.GetFriendApplyReq) (*pb_friend.GetFriendApplyResp, error) {
//	return nil, nil
//}

//func (rpc *RpcFriendServer) AddBlacklist(context.Context, *pb_friend.AddBlacklistReq) (*pb_ws.CommonResp, error) {
//	return nil, nil
//}
//func (rpc *RpcFriendServer) RemoveBlacklist(context.Context, *pb_friend.RemoveBlacklistReq) (*pb_ws.CommonResp, error) {
//	return nil, nil
//}
//func (rpc *RpcFriendServer) IsFriend(context.Context, *pb_friend.IsFriendReq) (*pb_friend.IsFriendResp, error) {
//	return nil, nil
//}
//func (rpc *RpcFriendServer) IsInBlackList(context.Context, *pb_friend.IsInBlackListReq) (*pb_friend.IsInBlackListResp, error) {
//	return nil, nil
//}
//func (rpc *RpcFriendServer) GetBlacklist(context.Context, *pb_friend.GetBlacklistReq) (*pb_friend.GetBlacklistResp, error) {
//	return nil, nil
//}

//func (rpc *RpcFriendServer) SetFriendComment(context.Context, *pb_friend.SetFriendCommentReq) (*pb_ws.CommonResp, error) {
//	return nil, nil
//}
//func (rpc *RpcFriendServer) ImportFriend(context.Context, *pb_friend.ImportFriendReq) (*pb_friend.ImportFriendResp, error) {
//	return nil, nil
//}
