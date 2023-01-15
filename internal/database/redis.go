package database

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type RedisDatabase struct {
	IDatabase // interface
	ctx       context.Context
	client    *redis.Client
}

func (rd *RedisDatabase) connect() {
	rd.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (rd *RedisDatabase) ConnectDatabase() *redis.Client {
	rd.connect()
	return rd.client
}

func NewRedisDatabase(ctx context.Context) *RedisDatabase {
	return &RedisDatabase{
		ctx: ctx,
	}
}
