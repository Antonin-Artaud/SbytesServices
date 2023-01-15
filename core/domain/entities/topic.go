package entities

import (
	"github.com/google/uuid"
	"time"
)

type Topic struct {
	Id        uuid.UUID `json:"id"`
	Subject   uuid.UUID `json:"subject"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}
