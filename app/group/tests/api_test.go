package tests

import (
	"context"
	"testing"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/pkg/common/discovery"
	"github.com/hashicorp/consul/api"
)

func TestGetGroups(t *testing.T) {
	consulCli, _ := api.NewClient(api.DefaultConfig())
	conn, err := discovery.GetGrpcConn(consulCli, "group-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewGroupClient(conn)

	resp, err := c.GetUserGroups(context.Background(), &pb.GetUserGroupsReq{UserId: 101})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	t.Log("3")
}

func TestGetGroupInfoById(t *testing.T) {
	consulCli, _ := api.NewClient(api.DefaultConfig())
	conn, err := discovery.GetGrpcConn(consulCli, "group-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewGroupClient(conn)

	resp, err := c.GetGroupInfoById(context.Background(), &pb.GetGroupInfoByIdReq{GroupId: 3})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	t.Log("3")
}

func TestGetGroupMember(t *testing.T) {
	consulCli, _ := api.NewClient(api.DefaultConfig())
	conn, err := discovery.GetGrpcConn(consulCli, "group-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewGroupClient(conn)

	resp, err := c.GetGroupMemberById(context.Background(), &pb.GetGroupMemberByIdReq{GroupId: 3})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	t.Log("3")
}

func TestCreateGroup(t *testing.T) {
	consulCli, _ := api.NewClient(api.DefaultConfig())
	conn, err := discovery.GetGrpcConn(consulCli, "group-srv")
	if err != nil {
		panic(err)
	}
	c := pb.NewGroupClient(conn)

	resp, err := c.CreateGroupInfo(context.Background(), &pb.CreateGroupInfoReq{
		GroupInfo: &pb.GroupInfo{
			GroupName:   "mygroup",
			CreatorId:   101,
			Description: "我的第一个群",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
