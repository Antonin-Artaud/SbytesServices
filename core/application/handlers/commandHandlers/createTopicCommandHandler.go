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

func (h *CreateTopicCommandHandler) Handle(ctx context.Context, command *commands.CreateTopicCommand) (*commandResponses.CreateTopicCommandResponse, error) {
	topic := &entities.Topic{
		Id:        uuid.New(),
		Subject:   command.Subject,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 5),
	}

	return &commandResponses.CreateTopicCommandResponse{
		Topic: topic,
	}, nil
}
