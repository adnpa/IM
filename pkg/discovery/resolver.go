package discovery

import (
	"context"
	"fmt"
	"github.com/adnpa/IM/pkg/common/config"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"strings"
	"sync"
	"time"
)

//https://github.com/grpc/grpc/blob/master/doc/naming.md
//参考dns_resolver和unix实现

var (
	nameResolver        = make(map[string]*Resolver)
	rwNameResolverMutex sync.RWMutex
)

func GetSrvConn(serviceName string) *grpc.ClientConn {
	return GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), serviceName)
}

// GetConn api
func GetConn(schema, etcdAddr, serviceName string) *grpc.ClientConn {
	rwNameResolverMutex.RLock()
	r, ok := nameResolver[schema+serviceName]
	if ok {
		rwNameResolverMutex.RUnlock()
		return r.grpcCli
	}
	r, err := newResolver(schema, etcdAddr, serviceName)
	if err != nil {
		rwNameResolverMutex.Unlock()
		return nil
	}
	rwNameResolverMutex.Unlock()
	return r.grpcCli
}

// ResolverBuilder 实现resolver.Builder接口
type ResolverBuilder struct {
	schema      string
	etcdAddr    string
	serviceName string
}

func NewBuilder() resolver.Builder {
	return &ResolverBuilder{}
}

// Build 创建并启动 etcd 解析器，用于监视目标的名称解析。
func (b *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r, err := newResolver(b.schema, b.etcdAddr, b.serviceName)
	if err != nil {
		return nil, err
	}
	r.cc = cc

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	key := GetPrefix(b.schema, r.serviceName)
	resp, err := r.etcdCli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var addrList []resolver.Address
	for _, ev := range resp.Kvs {
		addrList = append(addrList, resolver.Address{Addr: string(ev.Value)})
	}
	err = r.cc.UpdateState(resolver.State{Addresses: addrList})
	if err != nil {
		return nil, err
	}
	r.watchStartRevision = resp.Header.Revision + 1
	go r.watch(key, addrList)

	return r, nil
}

func (b *ResolverBuilder) Scheme() string {
	return b.schema
}

type Resolver struct {
	serviceName        string
	etcdAddr           string
	watchStartRevision int64

	cc      resolver.ClientConn
	grpcCli *grpc.ClientConn
	etcdCli *clientv3.Client
}

func newResolver(schema, etcdAddr, serviceName string) (*Resolver, error) {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(etcdAddr, ","),
	})
	if err != nil {
		return nil, err
	}

	grpcConn, err := grpc.NewClient(GetPrefix(schema, serviceName),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	r := &Resolver{
		//schema:      schema,
		serviceName: serviceName,
		etcdAddr:    etcdAddr,

		etcdCli: etcdCli,
		grpcCli: grpcConn,
	}

	return r, nil
}

func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {}

func (r *Resolver) Close() {}

func (r *Resolver) watch(prefix string, addrList []resolver.Address) {
	watchChan := r.etcdCli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for watchResp := range watchChan {
		changed := false
		for _, ev := range watchResp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				if !exists(addrList, string(ev.Kv.Value)) {
					changed = true
					addrList = append(addrList, resolver.Address{Addr: string(ev.Kv.Value)})
				}
			case mvccpb.DELETE:
				if s, ok := remove(addrList, string(ev.Kv.Value)); ok {
					changed = true
					addrList = s
				}
			}
		}

		if changed {
			err := r.cc.UpdateState(resolver.State{Addresses: addrList})
			if err != nil {
				log.Printf("update state err:%v", err)
			}
		}
	}
}

// helper
func exists(addrList []resolver.Address, addr string) bool {
	for _, v := range addrList {
		if v.Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}
