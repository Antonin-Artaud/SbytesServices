package server

import (
	"SbytesServices/core/application/controllers"
	"SbytesServices/internal/database"
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

type GinServerInterface interface {
	Configure()
	Run()
}

func NewGinServer(ctx context.Context, database database.Database, pLogger *logger.Logger, port int16) *GinServer {
	ginServer := &GinServer{
		engine:   nil,
		ctx:      ctx,
		database: database,
		logger:   pLogger,
		port:     port,
	}

	return ginServer
}

func (gs *GinServer) Configure() {
	gs.engine = gin.Default()
	err := gs.engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		gs.logger.LogError(err.Error())
	}
}

func (gs *GinServer) addControllers() {
	topicController := controllers.NewTopicController()

	routerBaseGroup := gs.engine.Group("/api/v1")
	{
		topicGroup := routerBaseGroup.Group("/topics")
		{
			topicGroup.POST("/", topicController.CreateTopic)
			topicGroup.GET("/:id", topicController.GetTopicById)
		}
	}
}

func (gs *GinServer) Run() {
	if gs.engine == nil {
		gs.Configure()
	}

	gs.addControllers()

	err := gs.engine.Run()
	if err != nil {
		panic(err)
	}
}
