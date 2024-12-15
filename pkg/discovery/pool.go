package discovery

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var (
	// ErrClosed is the error when the client pool is closed
	ErrClosed = errors.New("grpc pool: client pool is closed")
	// ErrTimeout is the error when the client pool timed out
	ErrTimeout = errors.New("grpc pool: client pool timed out")
	// ErrAlreadyClosed is the error when the client conn was already closed
	ErrAlreadyClosed = errors.New("grpc pool: the connection was already closed")
	// ErrFullPool is the error when the pool is already full
	ErrFullPool = errors.New("grpc pool: closing a ClientConn into a full pool")
)

type Factory func(schema, etcdaddr, servicename string) (*grpc.ClientConn, error)

type FactoryWithContext func(context.Context) (*grpc.ClientConn, error)

// Pool grpc连接池
type Pool struct {
	clients         chan ClientConn
	factoryFunc     FactoryWithContext
	idleTimeout     time.Duration
	maxLifeDuration time.Duration
	mu              sync.RWMutex
}

// ClientConn grpc conn wrapper
type ClientConn struct {
	*grpc.ClientConn
	pool          *Pool
	timeUsed      time.Time
	timeInitiated time.Time
	unhealthy     bool
}

func New(factory Factory, schema, etcdaddr, servicename string, init, capacity int, idleTimeout time.Duration, maxLifeDuration ...time.Duration) (*Pool, error) {
	factoryFunc := func(ctx context.Context) (*grpc.ClientConn, error) {
		return factory(schema, etcdaddr, servicename)
	}
	return NewWithContext(context.Background(), factoryFunc, init, capacity, idleTimeout, maxLifeDuration...)
}

func NewWithContext(ctx context.Context, factoryFunc FactoryWithContext, init, capacity int, idleTimeout time.Duration, maxLifeDuration ...time.Duration) (*Pool, error) {
	if capacity <= 0 {
		capacity = 1
	}
	if init < 0 {
		init = 0
	}
	if init > capacity {
		init = capacity
	}
	p := &Pool{
		clients:     make(chan ClientConn, capacity),
		factoryFunc: factoryFunc,
		idleTimeout: idleTimeout,
	}
	if len(maxLifeDuration) > 0 {
		p.maxLifeDuration = maxLifeDuration[0]
	}
	for i := 0; i < init; i++ {
		c, err := factoryFunc(ctx)
		if err != nil {
			return nil, err
		}
		p.clients <- ClientConn{
			ClientConn:    c,
			pool:          p,
			timeUsed:      time.Now(),
			timeInitiated: time.Now(),
		}
	}
	for i := 0; i < capacity-init; i++ {
		p.clients <- ClientConn{
			pool: p,
		}
	}
	return p, nil
}

func (p *Pool) IsClosed() bool {
	return p == nil || p.clients == nil
}

func (p *Pool) Close() {
	p.mu.Lock()
	clients := p.clients
	p.clients = nil
	p.mu.Unlock()

	if clients == nil {
		return
	}
	close(clients)

	for client := range clients {
		if client.ClientConn != nil {
			client.ClientConn.Close()
		}
	}
}

func (p *Pool) Get(ctx context.Context) (*ClientConn, error) {
	if p.IsClosed() {
		return nil, ErrClosed
	}

	wrapper := ClientConn{
		pool: p,
	}
	select {
	case wrapper = <-p.clients:
	case <-ctx.Done():
		return nil, ErrTimeout
	}

	// idle too long close & create
	if wrapper.ClientConn != nil && p.idleTimeout > 0 && wrapper.timeUsed.Add(p.idleTimeout).Before(time.Now()) {
		wrapper.ClientConn.Close()
		wrapper.ClientConn = nil
	}

	var err error
	if wrapper.ClientConn == nil {
		wrapper.ClientConn, err = p.factoryFunc(ctx)
		if err != nil {
			p.clients <- ClientConn{
				pool: p,
			}
		}
		wrapper.timeInitiated = time.Now()
		//todo wrapper.timeUsed = time.Now()
	}

	return &wrapper, nil
}
