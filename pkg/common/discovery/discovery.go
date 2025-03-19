package discovery

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetGrpcConn(consulConn *api.Client, name string) (*grpc.ClientConn, error) {
	srvInfo, err := GetHealthService(consulConn, name)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", srvInfo.Address, srvInfo.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	return conn, err
}

func GetHealthService(cli *api.Client, serviceName string) (*api.AgentService, error) {
	services, _, err := cli.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, fmt.Errorf("no healthy instances found for service: %s", serviceName)
	}

	return balanceService(services).Service, nil
}

// 负载均衡策略
func balanceService(services []*api.ServiceEntry) *api.ServiceEntry {
	return services[rand.Intn(len(services))]
}
