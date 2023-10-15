package events

import (
	"github.com/google/uuid"
)

type PlaceCreated struct {
	PlaceId uuid.UUID `json:"placeId" validate:"required"`
	Email   string    `json:"email" validate:"required"`
	EIN     string    `json:"ein" validate:"required"`
	Name    string    `json:"name" validate:"required"`
	Code    uuid.UUID `json:"code" validate:"required"` // code is used for signupcode in auth
}
