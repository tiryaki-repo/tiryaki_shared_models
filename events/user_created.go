package events

import (
	"time"

	"github.com/google/uuid"
)

type UserCreated struct {
	UserId         uuid.UUID  `json:"userId" validate:"required"`
	Email          string     `json:"email" validate:"required"`
	IsEmployee     bool       `json:"isEmployee" validate:"required"`
	BirthDate      time.Time  `json:"birthDate" validate:"required"`
	GivenName      string     `json:"givenName" validate:"required"`
	FamilyName     string     `json:"familyName" validate:"required"`
	PlaceId        *uuid.UUID `json:"placeId"`        // PlaceId is not needed for regular clients
	EmployeeTypeId *int       `json:"employeeTypeId"` // EmployeeTypeId is not needed for regular clients
}
