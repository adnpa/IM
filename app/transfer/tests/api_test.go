package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/adnpa/IM/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestSendMsg(t *testing.T) {
	conn, err := grpc.NewClient("192.168.8.37:50056", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPresenceClient(conn)
	resp, err := c.SendMsg(context.Background(), &pb.SendMsgReq{UserId: 1})
	// if err != nil {
	// 	t.Error(err)
	// }
	t.Log(err)
	fmt.Println(resp)
}
