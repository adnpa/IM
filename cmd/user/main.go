package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/adnpa/IM/internal/service/user"
	"github.com/adnpa/IM/pkg/common/pb"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	// *flag
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &user.UserService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// consul.Register(utils.Server)
}
