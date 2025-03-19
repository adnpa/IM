package storage

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"time"
)

// Pool maintains a pool of Redis connections.
type Pool interface {
	Get(ctx context.Context) (Conn, error)
}

// Conn is a single Redis connection.
type Conn interface {
	Exists(keys ...string) (int64, error)
	Get(name string) (string, error)
	Set(name string, value string) error
	SetNX(name string, value string, expiry time.Duration) error
	Incr(name string) (int64, error)
	Del(keys ...string) (int64, error)
	HGet(name, field string) (string, error)
	HGetAll(key string) (map[string]string, error)
	HSet(name string, value ...interface{}) error
	HDel(key string, fields ...string) (int64, error)
	HMGet(key string, fields ...string) ([]interface{}, error)
	HMSet(name string, values ...interface{}) error
	Eval(script *Script, keysAndArgs ...interface{}) (interface{}, error)
	PTTL(name string) (time.Duration, error)
	Close() error
}

// Script encapsulates the source, hash and key count for a Lua script.
// Taken from https://github.com/gomodule/redigo/blob/46992b0f02f74066bcdfd9b03e33bc03abd10dc7/redis/script.go#L24-L30
type Script struct {
	KeyCount int
	Src      string
	Hash     string
}

// NewScript returns a new script object. If keyCount is greater than or equal
// to zero, then the count is automatically inserted in the EVAL command
// argument list. If keyCount is less than zero, then the application supplies
// the count as the first value in the keysAndArgs argument to the Do, Send and
// SendHash methods.
// Taken from https://github.com/gomodule/redigo/blob/46992b0f02f74066bcdfd9b03e33bc03abd10dc7/redis/script.go#L32-L41
func NewScript(keyCount int, src string) *Script {
	h := sha1.New()
	_, _ = io.WriteString(h, src)
	return &Script{keyCount, src, hex.EncodeToString(h.Sum(nil))}
}
