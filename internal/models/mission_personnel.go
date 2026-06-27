package models

import (
	"time"

	"github.com/google/uuid"
)

type MissionPersonnel struct {
	ID         uuid.UUID  `json:"id"`
	MissionID  uuid.UUID  `json:"mission_id"`
	PeronnelID uuid.UUID  `json:"personnel_id"`
	Role       string     `json:"Role"`
	AssignedAt *time.Time `json:"assigned_at"`
}
