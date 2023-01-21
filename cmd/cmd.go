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

	pDatabase := database.NewRedisDatabase(pLogger)

	pDatabaseConnection, err := pDatabase.InitPoolConnections()

	if err != nil {
		pLogger.LogError(err.Error())
		return
	}

	defer closeDatabaseConnection(pDatabaseConnection, pLogger)

	registerMediatrHandler(pDatabase)

	createGinServer(ctx, pDatabase, pLogger)
}

func closeDatabaseConnection(pDatabase *redis.Client, pLogger *logger.Logger) {
	func(pClient *redis.Client) {
		err := pClient.Close()
		if err != nil {
			pLogger.LogError(err.Error())
		}
	}(pDatabase)
}

func registerMediatrHandler(databaseConnection database.IDatabase) {
	pCreateTopicCommandHandler := commandHandlers.NewCreateTopicCommandHandler(databaseConnection)
	pGetTopicByIdQueryHandler := queryHandlers.NewGetTopicByIdQueryHandler(databaseConnection)

	if err := mediatr.RegisterRequestHandler[*commands.CreateTopicCommand, *commandResponses.CreateTopicCommandResponse](pCreateTopicCommandHandler); err != nil {
		log.Fatalln(err)
	}

	if err := mediatr.RegisterRequestHandler[*queries.GetTopicByIdQuery, *queryResponses.GetTopicByIdQueryResponse](pGetTopicByIdQueryHandler); err != nil {
		log.Fatalln(err)
	}
}

func createGinServer(ctx context.Context, pDatabase *database.RedisDatabase, pLogger *logger.Logger) {
	pGinServer := server.NewGinServer(ctx, pDatabase, pLogger, 8080)
	pGinServer.Configure()
	pGinServer.Run()
}
