package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/pkg/common/discovery"
	"github.com/hashicorp/consul/api"
)

func TestGetOfflineMsg(t *testing.T) {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	conn, err := discovery.GetGrpcConn(consulCli, "offline-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewOfflineClient(conn)

	resp, err := c.GetOfflineMsg(context.Background(), &pb.GetOfflineMsgReq{Uid: 11})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp.Msgs)
}

func TestPutMsg(t *testing.T) {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	conn, err := discovery.GetGrpcConn(consulCli, "offline-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewOfflineClient(conn)

	resp, err := c.PutMsg(context.Background(), &pb.PutMsgReq{UserId: 11, Msg: &pb.ChatMsg{Content: "test content"}})
	t.Log(resp)
	t.Log(err)
	t.Log("c")
}

func TestRemoveMsg(t *testing.T) {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	conn, err := discovery.GetGrpcConn(consulCli, "offline-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewOfflineClient(conn)

	resp, err := c.RemoveMsg(context.Background(), &pb.RemoveMsgReq{MsgIds: []int64{0}})
	t.Log(resp)
	t.Log(err)
}
