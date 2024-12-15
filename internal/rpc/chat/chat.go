package chat

import (
	"github.com/adnpa/IM/common/config"
	"github.com/adnpa/IM/common/kafka"
	"github.com/adnpa/IM/discovery"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/pb"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
)

type RpcChatServer struct {
	*pb.UnimplementedChatServer

	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
	producer        *kafka.Producer
}

func NewRpcChatServer(port int) *RpcChatServer {
	rc := RpcChatServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.OfflineMessageName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
	rc.producer = kafka.NewProducer(config.Config.Kafka.Ws2mschat.Addr, config.Config.Kafka.Ws2mschat.Topic)
	return &rc
}

func (rpc *RpcChatServer) Run() {

	address := utils.ServerIP + ":" + strconv.Itoa(rpc.rpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return
	}

	//grpc server
	srv := grpc.NewServer()
	defer srv.GracefulStop()

	//service registers with etcd

	pb.RegisterChatServer(srv, rpc)
	err = discovery.Register(rpc.etcdSchema, strings.Join(rpc.etcdAddr, ","), utils.ServerIP, rpc.rpcPort, rpc.rpcRegisterName, 10)
	if err != nil {
		return
	}

	err = srv.Serve(listener)
	if err != nil {
		return
	}
}
