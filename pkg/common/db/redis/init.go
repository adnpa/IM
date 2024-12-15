package redis

import "github.com/redis/go-redis/v9"

var redisPool Pool

func init() {
	cli := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379", // use default Addr
			Password: "",               // no password set
			DB:       0,                // use default DB
		})
	redisPool = NewPool(cli)
}
