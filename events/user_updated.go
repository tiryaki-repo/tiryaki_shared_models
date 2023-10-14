package events

import (
	"github.com/google/uuid"
)

type UserUpdated struct {
	UserId     uuid.UUID `json:"userId" validate:"required"`
	GivenName  string    `json:"givenName" validate:"required"`
	FamilyName string    `json:"familyName" validate:"required"`
	Photo      *string   `json:"photo"`
}
