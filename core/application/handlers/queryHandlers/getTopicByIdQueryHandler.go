package queryHandlers

import (
	"SbytesServices/core/domain/entities"
	"SbytesServices/core/domain/queries"
	"SbytesServices/core/domain/queries/queryResponses"
	"SbytesServices/internal/database"
	"context"
)

type GetTopicByIdQueryHandler struct {
	database database.IDatabase
}

func NewGetTopicByIdQueryHandler(database database.IDatabase) *GetTopicByIdQueryHandler {
	return &GetTopicByIdQueryHandler{
		database: database,
	}
}

func (h *GetTopicByIdQueryHandler) Handle(_ context.Context, query *queries.GetTopicByIdQuery) (*queryResponses.GetTopicByIdQueryResponse, error) {
	if query.Id.String() == "00000000-0000-0000-0000-000000000000" {
		return &queryResponses.GetTopicByIdQueryResponse{
			Topic: &entities.Topic{
				Id: query.Id,
			},
		}, nil
	}

	return nil, nil
}
