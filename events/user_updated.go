package events

import (
	"github.com/google/uuid"
)

type UserUpdated struct {
	UserId     uuid.UUID `json:"userId"`
	GivenName  string    `json:"givenName"`
	FamilyName string    `json:"familyName"`
	Photo      *string   `json:"photo"`
}
