package controllers

import (
	"SbytesServices/core/domain/commands"
	"SbytesServices/core/domain/commands/commandResponses"
	"SbytesServices/core/domain/queries"
	"SbytesServices/core/domain/queries/queryResponses"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
)

type TopicController struct {
	IControllerBase // interface
}

func NewTopicController() *TopicController {
	return &TopicController{
		IControllerBase: &ControllerBase{},
	}
}

func (tc *TopicController) CreateTopic(ctx *gin.Context) {
	command := commands.CreateTopicCommand{}

	commandResponse, err := mediatr.Send[*commands.CreateTopicCommand, *commandResponses.CreateTopicCommandResponse](context.Background(), &command)

	if err != nil {
		tc.NotFound(ctx)
		return
	}

	tc.Created(ctx, commandResponse)
}

func (tc *TopicController) GetTopicById(ctx *gin.Context) {

	query := queries.GetTopicByIdQuery{}

	parse, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		tc.InternalServerError(ctx)
		return
	}

	query.Id = parse

	queryResponse, err := mediatr.Send[*queries.GetTopicByIdQuery, *queryResponses.GetTopicByIdQueryResponse](context.Background(), &query)

	if err != nil {
		tc.InternalServerError(ctx)
		return
	}

	if queryResponse.Topic == nil {
		tc.NotFound(ctx)
		return
	}

	tc.Found(ctx, queryResponse)
}
