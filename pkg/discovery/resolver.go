package discovery

import (
	"context"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/logger"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"strings"
	"sync"
	"time"
)

//https://github.com/grpc/grpc/blob/master/doc/naming.md
//参考dns_resolver和unix实现

var (
	//k-v本地缓存
	nameResolver        = make(map[string]*Resolver)
	rwNameResolverMutex sync.RWMutex
)

// GetConn api

func init() {

}

func GetSrvConn(serviceName string) *grpc.ClientConn {
	conn := GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), serviceName)
	return conn
}

func GetConn(schema, etcdAddr, serviceName string) *grpc.ClientConn {
	//todo  缓存conn
	b, _ := NewBuilder(schema, etcdAddr, serviceName)
	conn := NewGrpcConn()
	r, _ := b.Build(resolver.Target{}, conn, resolver.BuildOptions{})
	r.ResolveNow(resolver.ResolveNowOptions{})
	//
	//rwNameResolverMutex.RLock()
	//r, ok := nameResolver[schema+serviceName]
	//if ok {
	//	rwNameResolverMutex.RUnlock()
	//	return r.grpcCli
	//}
	//rwNameResolverMutex.RUnlock()
	//
	//rwNameResolverMutex.Lock()
	//r, ok = nameResolver[schema+serviceName]
	//
	//if ok {
	//	rwNameResolverMutex.Unlock()
	//	return r.grpcCli
	//}
	//
	//r, err := NewBuilder(schema, etcdAddr, serviceName).Build(resolver.Target{}, newEtcdDiscoveryMechanism(), resolver.BuildOptions{})
	//if err != nil {
	//	logger.L().Warn("build ", zap.Error(err))
	//	rwNameResolverMutex.Unlock()
	//	return nil
	//}
	//nameResolver[schema+serviceName] = r
	//rwNameResolverMutex.Unlock()
	return conn.Conn()
}

// ResolverBuilder 实现resolver.Builder接口
//type ResolverBuilder struct {
//	schema      string
//	etcdAddr    string
//	serviceName string
//}
//
//func NewBuilder(schema, etcdAddr, serviceName string) *ResolverBuilder {
//	return &ResolverBuilder{
//		schema:      schema,
//		etcdAddr:    etcdAddr,
//		serviceName: serviceName,
//	}
//}
//
//// Build 创建并启动 etcd 解析器，用于监视目标的名称解析。
//func (b *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (*Resolver, error) {
//	r, err := NewResolver(b.schema, b.etcdAddr, b.serviceName)
//	if err != nil {
//		return nil, err
//	}
//	r.cc = cc
//	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancelFunc()
//	key := GetPrefix(b.schema, r.serviceName)
//	resp, err := r.etcdCli.Get(ctx, key, clientv3.WithPrefix())
//	logger.L().Info("", zap.Any("resp", resp))
//	if err != nil {
//		logger.L().Warn("build resolver err", zap.Error(err))
//	}
//	var addrList []resolver.Address
//	for _, ev := range resp.Kvs {
//		addrList = append(addrList, resolver.Address{Addr: string(ev.Value)})
//	}
//	logger.L().Info("", zap.Any("addr", addrList))
//
//	err = r.cc.UpdateState(resolver.State{Addresses: addrList})
//
//	if err != nil {
//		logger.L().Warn("update err")
//		return nil, err
//	}
//	r.watchStartRevision = resp.Header.Revision + 1
//	go r.watch(key, addrList)
//
//	return r, nil
//}

//	func (b *ResolverBuilder) Scheme() string {
//		return b.schema
//	}
//
// uild(target Target, cc ClientConn, opts BuildOptions) (*Resolver, error)
func NewBuilder(schema, etcdAddr, serviceName string) (*Resolver, error) {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(etcdAddr, ","),
	})
	if err != nil {
		return nil, err
	}
	r := &Resolver{
		schema:      schema,
		etcdAddr:    etcdAddr,
		serviceName: serviceName,

		etcdCli: etcdCli,
	}
	return r, nil
}

type Resolver struct {
	//*manual.Resolver
	schema             string
	serviceName        string
	etcdAddr           string
	watchStartRevision int64

	cc      resolver.ClientConn
	grpcCli *grpc.ClientConn
	etcdCli *clientv3.Client

	//r.lastSeenState = &s
}

//func (r *Resolver) Conn() {
//	r.cc.
//}

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r, err := NewBuilder(r.schema, r.etcdAddr, r.serviceName)
	if err != nil {
		return nil, err
	}
	r.cc = cc

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	key := GetPrefix(r.schema, r.serviceName)
	resp, err := r.etcdCli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		logger.L().Warn("build resolver err", zap.Error(err))
	}
	var addrList []resolver.Address
	for _, ev := range resp.Kvs {
		addrList = append(addrList, resolver.Address{Addr: string(ev.Value)})
	}
	err = r.cc.UpdateState(resolver.State{Addresses: addrList})
	if err != nil {
		logger.L().Warn("update err")
		return nil, err
	}

	r.watchStartRevision = resp.Header.Revision + 1
	go r.watch(key, addrList)

	return r, nil
}

func (r *Resolver) Scheme() string {
	return r.schema
}

func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {
	//r.cc.UpdateState()
}

func (r *Resolver) Close() {}

func (r *Resolver) UpdateState(s resolver.State) {

}
func (r *Resolver) ReportError(err error) {

}

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
				logger.L().Warn("service update err", zap.Error(err))
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
