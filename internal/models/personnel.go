package models

import (
	"time"

	"github.com/google/uuid"
)

type Personnel struct {
	ID             uuid.UUID  `json:"id"`
	Rank           string     `json:"rank"`
	LastName       string     `json:"last_name"`
	FirstName      string     `json:"first_name"`
	UnitDesignator string     `json:"unit_designator"`
	ClearanceLevel string     `json:"clearance_level"`
	Status         string     `json:"status"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
