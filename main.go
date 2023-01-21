package main

import (
	"SbytesServices/internal/database"
	"SbytesServices/internal/helpers/logger"
)

func main() {
	//cmd.Start()

	db := database.NewRedisDatabase(logger.NewLogger(logger.INFO))

	// Init pool connections
	_, err := db.InitPoolConnections()

	if err != nil {
		panic(err)
	}

	// Ping
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	// Write entity
	err = db.WriteEntity("key", "value")

	if err != nil {
		panic(err)
	}

	err = db.RedisPool.Close()
	if err != nil {
		return
	}
}
