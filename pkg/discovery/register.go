package discovery

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"net"
	"strconv"
	"strings"
)

//key: schema:///serviceName/

var reg *Reg

type Reg struct {
	cli    *clientv3.Client
	ctx    context.Context
	cancel context.CancelFunc
	key    string
}

func Register(schema, etcdAddr, host string, port int, serviceName string, ttl int) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(etcdAddr, ","),
	})
	if err != nil {
		log.Println(err)
	}

	//lease
	ctx, cancelFunc := context.WithCancel(context.Background())
	resp, err := cli.Grant(ctx, int64(ttl))
	if err != nil {
		log.Println(err)
	}

	//put
	val := net.JoinHostPort(host, strconv.Itoa(port))
	key := GetPrefix(schema, serviceName) + val
	_, err = cli.Put(ctx, key, val)

	//keepalive
	keepAliveResponses, err := cli.KeepAlive(ctx, resp.ID)
	if err != nil {
		log.Println(err)
	}

	//heartbeat
	go func() {
		for {
			select {
			case _, ok := <-keepAliveResponses:
				if !ok {
					return
				}
			}
		}
	}()

	reg = &Reg{
		cli:    cli,
		ctx:    ctx,
		cancel: cancelFunc,
		key:    key,
	}

	return nil
}

func GetPrefix(schema string, name string) string {
	return fmt.Sprintf("%s:///%s", schema, name)
}

func GetPrefix4Unique(schema, serviceName string) string {
	return fmt.Sprintf("%s:///%s", schema, serviceName)
}

func RegisterUnique(schema, etcdAddr, myHost string, myPort int, serviceName string, ttl int) error {
	serviceName = serviceName + ":" + net.JoinHostPort(myHost, strconv.Itoa(myPort))
	return Register(schema, etcdAddr, myHost, myPort, serviceName, ttl)
}

func Unregister() {
	reg.cancel()
	reg.cli.Delete(reg.ctx, reg.key)
}