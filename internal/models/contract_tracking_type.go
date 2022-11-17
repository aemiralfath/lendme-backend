package models

import (
	"github.com/google/uuid"
	"time"
)

type ContractTrackingType struct {
	ContractTrackingID uuid.UUID `json:"contract_tracking_id" db:"contract_tracking_id" binding:"omitempty"`
	Name               string    `json:"name" db:"name" binding:"omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
