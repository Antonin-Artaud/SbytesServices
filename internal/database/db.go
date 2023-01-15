package database

import "github.com/go-redis/redis/v9"

type Database interface {
	ConnectDatabase() *redis.Client
}
