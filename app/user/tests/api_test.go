package tests

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/adnpa/IM/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetUser(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)
	resp, err := c.GetUserById(context.Background(), &pb.GetUserByIdReq{Id: 101})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp.Usr)
}

func TestCreateUser(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)
	resp, err := c.CreateUser(context.Background(), &pb.CreateUserReq{
		Nickname: "dddd",
		Mobile:   "444444",
		Email:    "dddd@gmail.com",
		Password: "111111",
	})
	t.Log(resp)
	t.Log(err)
	t.Log()
}

func TestMatch(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)

	u, _ := c.GetUserById(context.Background(), &pb.GetUserByIdReq{Id: 101})

	resp, err := c.CheckPassWord(context.Background(), &pb.CheckPassWordReq{
		Password:          "111111",
		EncryptedPassword: u.Usr.PassWord,
		Salt:              u.Usr.Salt,
	})
	t.Log(resp.Match)
	t.Log(err)
}
