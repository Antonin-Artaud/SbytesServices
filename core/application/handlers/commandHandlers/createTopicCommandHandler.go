package commandHandlers

import (
	"SbytesServices/core/domain/commands"
	"SbytesServices/core/domain/commands/commandResponses"
	"SbytesServices/core/domain/entities"
	"context"
	"github.com/google/uuid"
	"time"
)

type CreateTopicCommandHandler struct {
}

func NewCreateTopicCommandHandler() *CreateTopicCommandHandler {
	return &CreateTopicCommandHandler{}
}

func (h *CreateTopicCommandHandler) Handle(_ context.Context, _ *commands.CreateTopicCommand) (*commandResponses.CreateTopicCommandResponse, error) {
	var (
		topicId uuid.UUID
		subject uuid.UUID
		err     error
	)

	if topicId, err = uuid.NewRandom(); err != nil {
		return nil, err
	}

	if subject, err = uuid.NewRandom(); err != nil {
		return nil, err
	}

	response := &commandResponses.CreateTopicCommandResponse{
		Topic: &entities.Topic{
			Id:        topicId,
			Subject:   subject,
			CreatedAt: time.Now(),
			ExpiredAt: time.Now().Add(time.Minute * 5),
		},
	}

	return response, nil
}
