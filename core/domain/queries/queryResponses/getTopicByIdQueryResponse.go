package queryResponses

import "SbytesServices/core/domain/entities"

type GetTopicByIdQueryResponse struct {
	Topic *entities.Topic `json:"topic"`
}
