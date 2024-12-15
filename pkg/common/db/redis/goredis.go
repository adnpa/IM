package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type pool struct {
	delegate *redis.Client
}

func (p *pool) Get(ctx context.Context) (Conn, error) {
	c := p.delegate
	return &conn{delegate: c, ctx: ctx}, nil
}

// NewPool returns a Goredis-based pool implementation.
func NewPool(delegate *redis.Client) Pool {
	return &pool{delegate}
}

type conn struct {
	delegate *redis.Client
	ctx      context.Context
}

func (c *conn) Get(name string) (string, error) {
	value, err := c.delegate.Get(c.ctx, name).Result()
	return value, noErrNil(err)
}

func (c *conn) Set(name string, value string) (bool, error) {
	reply, err := c.delegate.Set(c.ctx, name, value, 0).Result()
	return reply == "OK", noErrNil(err)
}

func (c *conn) SetNX(name string, value string, expiry time.Duration) (bool, error) {
	ok, err := c.delegate.SetNX(c.ctx, name, value, expiry).Result()
	return ok, noErrNil(err)
}

func (c *conn) PTTL(name string) (time.Duration, error) {
	expiry, err := c.delegate.PTTL(c.ctx, name).Result()
	return expiry, noErrNil(err)
}

func (c *conn) Eval(script *Script, keysAndArgs ...interface{}) (interface{}, error) {
	keys := make([]string, script.KeyCount)
	args := keysAndArgs

	if script.KeyCount > 0 {
		for i := 0; i < script.KeyCount; i++ {
			keys[i] = keysAndArgs[i].(string)
		}

		args = keysAndArgs[script.KeyCount:]
	}

	v, err := c.delegate.EvalSha(c.ctx, script.Hash, keys, args...).Result()
	if err != nil && strings.HasPrefix(err.Error(), "NOSCRIPT ") {
		v, err = c.delegate.Eval(c.ctx, script.Src, keys, args...).Result()
	}
	return v, noErrNil(err)
}

func (c *conn) Incr(name string) (int64, error) {
	result, err := c.delegate.Incr(c.ctx, name).Result()
	return result, err
}

func (c *conn) Close() error {
	return c.delegate.Close()
}

func noErrNil(err error) error {
	if !errors.Is(err, redis.Nil) {
		return err
	}
	return nil
}