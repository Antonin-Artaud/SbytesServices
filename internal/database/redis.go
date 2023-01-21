package database

import (
	"SbytesServices/internal/helpers/logger"
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
)

type RedisDatabase struct {
	IDatabase // interface
	ctx       context.Context
	client    *redis.Client
	logger    *logger.Logger
}

func (rd *RedisDatabase) connect() error {
	rd.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pingResponse, err := rd.client.Ping(rd.ctx).Result()

	rd.logger.LogInfo(pingResponse)

	if err != nil {
		return err
	}

	return nil
}

func (rd *RedisDatabase) ConnectDatabase() (*redis.Client, error) {
	if err := rd.connect(); err != nil {
		return nil, err
	}

	return rd.client, nil
}

func NewRedisDatabase(ctx context.Context, logger *logger.Logger) *RedisDatabase {
	return &RedisDatabase{
		ctx:    ctx,
		logger: logger,
	}
}

func (rd *RedisDatabase) GetConnection() (*redis.Client, error) {
	if rd.client == nil {
		return nil, errors.New("redis client is nil")
	}

	return rd.client, nil
}
