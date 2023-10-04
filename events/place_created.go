package events

import (
	"github.com/google/uuid"
)

type PlaceCreated struct {
	PlaceId uuid.UUID `json:"placeId"`
	EIN     string    `json:"ein"`
	Name    string    `json:"name"`
	Code    uuid.UUID `json:"code"`
}
