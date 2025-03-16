package service_test

import (
	"context"
	"log"
	"testing"

	"github.com/adnpa/IM/pkg/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetUser(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)
	resp, _ := c.GetUserById(context.Background(), &pb.GetUserByIdReq{Id: 0})
	t.Log(resp.Usr.PassWord)
	t.Log(resp.Usr.Salt)
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
