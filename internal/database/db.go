package database

import "github.com/go-redis/redis/v9"

// IDatabase interface
type IDatabase interface {
	ConnectDatabase() (*redis.Client, error)
}
