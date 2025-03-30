package tests

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/pkg/common/discovery"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetFriends(t *testing.T) {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	conn, err := discovery.GetGrpcConn(consulCli, "friend-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewFriendClient(conn)
	resp, err := c.GetFriendsByUserId(context.Background(), &pb.GetFriendsByUserIdReq{Uid: 101})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
	t.Log(err)
	t.Log("")
}

func TestGetFriendApplys(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFriendClient(conn)
	resp, err := c.GetFriendApply(context.Background(), &pb.GetFriendApplyReq{UserId: 101})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp.FriendApplys)
	t.Log(err)

}

func TestCreateUser(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)
	c.CreateUser(context.Background(), &pb.CreateUserReq{})
}

func TestMatch(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)

	u, _ := c.GetUserById(context.Background(), &pb.GetUserByIdReq{Id: 0})

	resp, _ := c.CheckPassWord(context.Background(), &pb.CheckPassWordReq{
		Password:          "",
		EncryptedPassword: u.Usr.PassWord,
		Salt:              u.Usr.Salt,
	})
	t.Log(resp.Match)
}
