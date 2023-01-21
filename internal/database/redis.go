package database

import (
	"SbytesServices/internal/helpers/logger"
	"github.com/mediocregopher/radix/v3"
)

type RedisDatabase struct {
	IDatabase // interface
	RedisPool *radix.Pool
	logger    *logger.Logger
}

func (rd *RedisDatabase) InitPoolConnections() (*radix.Pool, error) {
	pool, err := radix.NewPool("tcp", "localhost:6379", 10)

	if err != nil {
		rd.logger.LogError(err.Error())
		return nil, err
	}

	rd.RedisPool = pool

	return pool, nil
}

func (rd *RedisDatabase) Ping() error {
	err := rd.RedisPool.Do(radix.Cmd(nil, "PING"))

	if err != nil {
		rd.logger.LogError(err.Error())
		return err
	}

	return nil
}

func (rd *RedisDatabase) WriteEntity(key string, value string) error {

	err := rd.RedisPool.Do(radix.Cmd(nil, "MULTI"))
	if err != nil {
		rd.logger.LogError(err.Error())
		return err
	}

	err = rd.RedisPool.Do(radix.Cmd(nil, "SET", key, value))
	if err != nil {
		rd.logger.LogError(err.Error())
		return err
	}

	err = rd.RedisPool.Do(radix.Cmd(nil, "EXEC"))
	if err != nil {
		rd.logger.LogError(err.Error())
		return err
	}

	return nil
}

func (rd *RedisDatabase) ReadEntity(key string) (string, error) {
	var value string
	err := rd.RedisPool.Do(radix.Cmd(&value, "GET", key))
	if err != nil {
		rd.logger.LogError(err.Error())
		return "", err
	}

	return value, nil
}

func NewRedisDatabase(logger *logger.Logger) *RedisDatabase {
	return &RedisDatabase{
		RedisPool: nil,
		logger:    logger,
	}
}
