package queryHandlers

import (
	"SbytesServices/core/domain/queries"
	"SbytesServices/core/domain/queries/queryResponses"
	"context"
)

type GetTopicByIdQueryHandler struct {
}

func NewGetTopicByIdQueryHandler() *GetTopicByIdQueryHandler {
	return &GetTopicByIdQueryHandler{}
}

func (h *GetTopicByIdQueryHandler) Handle(ctx context.Context, query *queries.GetTopicByIdQuery) (*queryResponses.GetTopicByIdQueryResponse, error) {
	return &queryResponses.GetTopicByIdQueryResponse{
		Topic: nil,
	}, nil
}
