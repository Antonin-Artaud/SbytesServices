package cmd

import (
	"SbytesServices/core/application/handlers/commandHandlers"
	"SbytesServices/core/application/handlers/queryHandlers"
	"SbytesServices/core/domain/commands"
	"SbytesServices/core/domain/commands/commandResponses"
	"SbytesServices/core/domain/queries"
	"SbytesServices/core/domain/queries/queryResponses"
	"SbytesServices/internal/database"
	"SbytesServices/internal/helpers/logger"
	"SbytesServices/server"
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/mehdihadeli/go-mediatr"
	"log"
)

func Start() {
	ctx := context.Background()

	pLogger := logger.NewLogger(logger.INFO)

	pRedisDatabase := database.NewRedisDatabase(ctx)

	pDatabase := pRedisDatabase.ConnectDatabase()

	defer func(pClient *redis.Client) {
		err := pClient.Close()
		if err != nil {
			pLogger.LogError(err.Error())
		}
	}(pDatabase)

	registerMediatrHandler()

	ginServer := server.NewGinServer(ctx, pRedisDatabase, pLogger, 8080)
	ginServer.Configure()
	ginServer.Run()
}

func registerMediatrHandler() {
	createTopicCommandHandler := commandHandlers.NewCreateTopicCommandHandler()
	getTopicByIdQueryHandler := queryHandlers.NewGetTopicByIdQueryHandler()

	if err := mediatr.RegisterRequestHandler[*commands.CreateTopicCommand, *commandResponses.CreateTopicCommandResponse](createTopicCommandHandler); err != nil {
		log.Fatalln(err)
	}

	if err := mediatr.RegisterRequestHandler[*queries.GetTopicByIdQuery, *queryResponses.GetTopicByIdQueryResponse](getTopicByIdQueryHandler); err != nil {
		log.Fatalln(err)
	}
}
