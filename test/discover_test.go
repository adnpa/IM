package test

import (
	"context"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/discovery"
	"github.com/adnpa/IM/pkg/pb/pb_user"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
)

func TestEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		//Endpoints: append(make([]string, 0), "192.168.1.129:2379"),
		Endpoints: []string{"192.168.1.129:2379"},
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	putResp, err := cli.Put(ctx, "testKey", "testValue")
	cli.Put(ctx, "testKey", "testValue1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(putResp)

	getResp, err := cli.Get(ctx, "testKey")
	if err != nil {
		t.Fatal(err)
	}
	for _, ev := range getResp.Kvs {
		t.Log(string(ev.Value))
		t.Log(string(ev.Key))
	}
	t.Log("get")

}

func TestResolver(t *testing.T) {
	//err := discovery.Register("goIM", "192.168.1.129:2379", "127.0.0.1", config.Config.RpcPort.UserPort[0], config.Config.RpcRegisterName.UserName, 10)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//t.Log("register success")

	//conn := discovery.GetSrvConn("testName")
	conn := discovery.GetSrvConn(config.Config.RpcRegisterName.UserName)
	t.Log("conn", conn)
	client := pb_user.NewUserClient(conn)
	t.Log("cli", client)

	pbParams := &pb_user.GetUserInfoReq{}
	_, err := client.GetUserInfo(context.Background(), pbParams)
	t.Log("err", err)
	//select {}
}
