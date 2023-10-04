package events

import (
	"github.com/google/uuid"
)

type PlaceCreated struct {
	PlaceId uuid.UUID
	EIN     string
	Name    string
	Code    uuid.UUID
}
