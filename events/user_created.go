package events

import (
	"time"

	"github.com/google/uuid"
)

type UserCreated struct {
	UserId         uuid.UUID  `json:"userId"`
	Email          string     `json:"email"`
	IsEmployee     bool       `json:"isEmployee"`
	BirthDate      time.Time  `json:"birthDate"`
	GivenName      string     `json:"givenName"`
	FamilyName     string     `json:"familyName"`
	PlaceId        *uuid.UUID `json:"placeId"`        // PlaceId is not needed for regular clients
	EmployeeTypeId *int       `json:"employeeTypeId"` // EmployeeTypeId is not needed for regular clients
}
