package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/redis/go-redis/v9"
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
	t.Log(err)
	fmt.Println(resp)
}

func TestIsOnline(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50056", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPresenceClient(conn)
	resp, err := c.IsOnline(context.Background(), &pb.IsOnlineReq{UserId: 1})
	t.Log(err)
	fmt.Println(resp)
}

func TestGetAllLoginUser(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		Protocol: 3,  // specify 2 for RESP 2 or 3 for RESP 3
	})
	resp, err := rdb.Keys(context.Background(), "*").Result()
	t.Log(err)
	for _, user_id := range resp {
		t.Log("user", user_id)
		server := rdb.Get(context.Background(), user_id)
		logger.Infof("login user", "user", user_id, "server", server.Val())
	}
}
