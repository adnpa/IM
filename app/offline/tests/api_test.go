package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/adnpa/IM/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetOfflineMsg(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewOfflineClient(conn)
	resp, err := c.GetOfflineMsg(context.Background(), &pb.GetOfflineMsgReq{Uid: 1})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp.Msgs)
}

func TestPutMsg(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewOfflineClient(conn)
	resp, err := c.PutMsg(context.Background(), &pb.PutMsgReq{Msg: &pb.ChatMsg{Content: "test content"}})
	t.Log(resp)
	t.Log(err)
	t.Log()
}

func TestRemoveMsg(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewOfflineClient(conn)
	resp, err := c.RemoveMsg(context.Background(), &pb.RemoveMsgReq{MsgIds: []int64{0}})
	t.Log(resp)
	t.Log(err)
}
