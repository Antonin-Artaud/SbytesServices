package server

import (
	"SbytesServices/internal/helpers/database"
	"SbytesServices/internal/helpers/logger"
	"context"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	engine   *gin.Engine
	ctx      context.Context
	database database.Database
	logger   *logger.Logger
	port     int16
}

func NewGinServer(ctx context.Context, database database.Database, pLogger *logger.Logger, port int16) *GinServer {
	ginServer := &GinServer{
		engine:   gin.Default(),
		ctx:      ctx,
		database: database,
		logger:   pLogger,
		port:     port,
	}

	return ginServer
}

func (gs *GinServer) Run() {
	err := gs.engine.Run()
	if err != nil {
		panic(err)
	}
}
