package commandResponses

import "SbytesServices/core/domain/entities"

type CreateTopicCommandResponse struct {
	Topic *entities.Topic `json:"topic"`
}
