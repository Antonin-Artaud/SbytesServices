package database

import "github.com/go-redis/redis/v9"

// IDatabase interface
type IDatabase interface {
	ConnectDatabase() (*redis.Client, error)
	WriteEntity(key string, value string) error
	ReadEntity(key string) (string, error)
}
