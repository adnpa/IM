package discovery

import (
	"fmt"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

type GrpcConn struct {
	conn  *grpc.ClientConn
	state resolver.State
}

func NewGrpcConn() *GrpcConn {
	return &GrpcConn{}
}

func (c *GrpcConn) Conn() *grpc.ClientConn {
	return c.conn
}

func (c *GrpcConn) UpdateState(state resolver.State) error {
	c.state = state
	if len(state.Addresses) <= 0 {
		logger.L().Info("new addr is null")
		return nil
	}

	addresses := c.state.Addresses
	idx := utils.RandIntN(len(addresses))
	if c.conn != nil {
		c.conn.Close()
	}
	newConn, err := grpc.NewClient(addresses[idx].Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.L().Info("err", zap.Error(err))
		return err
	}
	c.conn = newConn
	return nil
}

func (c *GrpcConn) ReportError(error) {
}

func (c *GrpcConn) NewAddress(addresses []resolver.Address) {

}

func (c *GrpcConn) ParseServiceConfig(serviceConfigJSON string) *serviceconfig.ParseResult {
	return &serviceconfig.ParseResult{Err: fmt.Errorf("service config not supported")}
}
