package commands

import "time"

type CreateTopicCommand struct {
	Subject string `json:"subject"`
}

type CreateTopicCommandResponse struct {
	Id        string    `json:"id"`
	Subject   string    `json:"subject"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}
