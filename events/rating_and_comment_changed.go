package events

import (
	"time"

	"github.com/google/uuid"
)

type RatingAndCommentStatus int

const (
	Deleted  RatingAndCommentStatus = 0
	Inserted RatingAndCommentStatus = 1
)

type RatingAndCommentChanged struct {
	Id        uuid.UUID              `json:"id" validate:"required"`
	PlaceId   uuid.UUID              `json:"placeId" validate:"required"`
	ClientId  uuid.UUID              `json:"clientId" validate:"required"`
	Rating    int                    `json:"rating" validate:"required"`
	Comment   *string                `json:"comment"`
	CreatedAt time.Time              `json:"createdAt" validate:"required"`
	Status    RatingAndCommentStatus `json:"status" validate:"required"`
}
