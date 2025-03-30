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

func TestGetUser(t *testing.T) {
	consulCli, _ := api.NewClient(api.DefaultConfig())
	conn, err := discovery.GetGrpcConn(consulCli, "user-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewUserClient(conn)

	resp, err := c.GetUserById(context.Background(), &pb.GetUserByIdReq{Id: 101})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp.Usr)
	t.Log(resp.Usr.Sex)
	t.Log(resp.Usr.Memo)
	t.Log("ff")
}

func TestCreateUser(t *testing.T) {
	consulCli, _ := api.NewClient(api.DefaultConfig())
	conn, err := discovery.GetGrpcConn(consulCli, "user-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewUserClient(conn)

	c.CreateUser(context.Background(), &pb.CreateUserReq{})
}

func TestUpdateUser(t *testing.T) {
	consulCli, _ := api.NewClient(api.DefaultConfig())
	conn, err := discovery.GetGrpcConn(consulCli, "user-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewUserClient(conn)

	c.UpdateUser(context.Background(), &pb.UpdateUserReq{
		Id:       101,
		Nickname: "new name",
	})
	t.Log("22")
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
