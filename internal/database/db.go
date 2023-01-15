package database

import "github.com/go-redis/redis/v9"

type IDatabase interface {
	ConnectDatabase() *redis.Client
}
