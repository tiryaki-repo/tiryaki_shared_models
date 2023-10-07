package events

import (
	"github.com/google/uuid"
)

type UserJoinRequested struct {
	RequesterUserID uuid.UUID `json:"requesterUserId"`
	RequesteeUserId uuid.UUID `json:"requesteeUserId"`
	PlaceId         uuid.UUID `json:"placeId"`
	EmployeeTypeId  int       `json:"employeeTypeId"`
	Code            uuid.UUID `json:"code"` // code is used for signupcode in auth
}
