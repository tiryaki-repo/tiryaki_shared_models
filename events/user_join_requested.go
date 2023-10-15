package events

import (
	"github.com/google/uuid"
)

type UserJoinRequested struct {
	RequesterUserId    uuid.UUID  `json:"requesterUserId" validate:"required"`
	Email              string     `json:"email" validate:"required"`
	RequesteeUserId    *uuid.UUID `json:"requesteeUserId"`
	RequesteeFirstName *string    `json:"RequesteeFirstName"`
	PlaceId            uuid.UUID  `json:"placeId" validate:"required"`
	EmployeeTypeId     int        `json:"employeeTypeId" validate:"required"`
	Code               uuid.UUID  `json:"code" validate:"required"` // code is used for signupcode in auth
}
