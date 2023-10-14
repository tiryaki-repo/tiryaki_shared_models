package events

import (
	"github.com/google/uuid"
)

type UserJoinRequested struct {
	RequesterUserID uuid.UUID `json:"requesterUserId" validate:"required"`
	Email           string    `json:"requesteeUserId" validate:"required"`
	PlaceId         uuid.UUID `json:"placeId" validate:"required"`
	EmployeeTypeId  int       `json:"employeeTypeId" validate:"required"`
	Code            uuid.UUID `json:"code" validate:"required"` // code is used for signupcode in auth
}
