package database

import (
	"github.com/mediocregopher/radix/v3"
)

// IDatabase interface
type IDatabase interface {
	InitPoolConnections() (*radix.Pool, error)
	Ping() error
	WriteEntity(key string, value string) error
	ReadEntity(key string) (string, error)
}
