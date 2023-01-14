package queries

import "github.com/google/uuid"

type GetTopicByIdQuery struct {
	Id uuid.UUID `json:"id"`
}
