package gateway

import (
	"context"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/adnpa/IM/pkg/discovery"
	"github.com/adnpa/IM/pkg/pb/pb_relay"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"strings"
)

type RpcRelayServer struct {
	*pb_relay.UnimplementedOnlineMessageRelayServiceServer

	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func (s *RpcRelayServer) Init(rpcPort int) {
	s.rpcPort = rpcPort
	s.rpcRegisterName = config.Config.RpcRegisterName.OnlineMessageRelayName
	s.etcdSchema = config.Config.Etcd.EtcdSchema
	s.etcdAddr = config.Config.Etcd.EtcdAddr
}
func (s *RpcRelayServer) Run() {
	ip := utils.ServerIP
	registerAddress := ip + ":" + utils.IntToString(s.rpcPort)
	listener, err := net.Listen("tcp", registerAddress)
	if err != nil {
		return
	}
	defer listener.Close()
	srv := grpc.NewServer()
	defer srv.GracefulStop()
	pb_relay.RegisterOnlineMessageRelayServiceServer(srv, s)
	err = discovery.RegisterUnique(s.etcdSchema, strings.Join(s.etcdAddr, ","), ip, s.rpcPort, s.rpcRegisterName, 10)
	if err != nil {
		log.Println(err)
	}
	err = srv.Serve(listener)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *RpcRelayServer) OnlinePushMsg(ctx context.Context, in *pb_relay.OnlinePushMsgReq) (*pb_relay.OnlinePushMsgResp, error) {
	var resp []*pb_relay.SingleMsgToUser
	msgBytes, _ := proto.Marshal(in.MsgData)
	mReply := Resp{
		ReqIdentifier: constant.WSPushMsg,
		OperationID:   in.OperationID,
		Data:          msgBytes,
	}
	replyBytes, err := utils.MarshalGob(mReply)
	if err != nil {
		return nil, err
	}

	recvID := in.PushToUserID
	platformList := genPlatformArray()
	for _, v := range platformList {
		if conn := ws.getWsConn(recvID, v); conn != nil {
			resultCode := sendMsgToUser(conn, replyBytes, in, v, recvID)
			temp := &pb_relay.SingleMsgToUser{
				ResultCode:     resultCode,
				RecvID:         recvID,
				RecvPlatFormID: constant.PlatformNameToID(v),
			}
			resp = append(resp, temp)
		} else {
			temp := &pb_relay.SingleMsgToUser{
				ResultCode:     -1,
				RecvID:         recvID,
				RecvPlatFormID: constant.PlatformNameToID(v),
			}
			resp = append(resp, temp)
		}
	}
	return &pb_relay.OnlinePushMsgResp{
		Resp: resp,
	}, nil
}

func (s *RpcRelayServer) GetUsersOnlineStatus(ctx context.Context, req *pb_relay.GetUsersOnlineStatusReq) (*pb_relay.GetUsersOnlineStatusResp, error) {
	var resp pb_relay.GetUsersOnlineStatusResp
	for _, userID := range req.UserIDList {
		platformList := genPlatformArray()
		temp := new(pb_relay.GetUsersOnlineStatusResp_SuccessResult)
		temp.UserID = userID
		for _, platform := range platformList {
			if conn := ws.getWsConn(userID, platform); conn != nil {
				ps := new(pb_relay.GetUsersOnlineStatusResp_SuccessDetail)
				ps.Platform = platform
				ps.Status = constant.OnlineStatus
				temp.Status = constant.OnlineStatus
				temp.DetailPlatformStatus = append(temp.DetailPlatformStatus, ps)

			}
		}
		if temp.Status == constant.OnlineStatus {
			resp.SuccessResult = append(resp.SuccessResult, temp)
		}
	}
	return &resp, nil
}

func sendMsgToUser(conn *WsConn, bMsg []byte, in *pb_relay.OnlinePushMsgReq, RecvPlatForm, RecvID string) (ResultCode int64) {
	err := ws.writeMsg(conn, websocket.BinaryMessage, bMsg)
	if err != nil {
		ResultCode = -2
		return ResultCode
	} else {
		ResultCode = 0
		return ResultCode
	}
}

func genPlatformArray() (array []string) {
	for i := 1; i <= constant.LinuxPlatformID; i++ {
		array = append(array, constant.PlatformIDToName(int32(i)))
	}
	return array
}
