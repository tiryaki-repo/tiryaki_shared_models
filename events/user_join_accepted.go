package events

import (
	"github.com/google/uuid"
)

type UserJoinAccepted struct {
	UserId         uuid.UUID `json:"userId" validate:"required"`
	Email          string    `json:"email" validate:"required"`
	GivenName      string    `json:"givenName" validate:"required"`
	FamilyName     string    `json:"familyName" validate:"required"`
	PlaceId        uuid.UUID `json:"placeId" validate:"required"`
	EmployeeTypeId int       `json:"employeeTypeId" validate:"required"`
}
