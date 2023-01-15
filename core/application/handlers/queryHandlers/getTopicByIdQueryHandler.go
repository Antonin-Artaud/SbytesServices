package queryHandlers

import (
	"SbytesServices/core/domain/entities"
	"SbytesServices/core/domain/queries"
	"SbytesServices/core/domain/queries/queryResponses"
	"context"
)

type GetTopicByIdQueryHandler struct {
}

func NewGetTopicByIdQueryHandler() *GetTopicByIdQueryHandler {
	return &GetTopicByIdQueryHandler{}
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
