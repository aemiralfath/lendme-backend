package models

import "time"

type InstallmentStatusType struct {
	InstallmentStatusID int       `json:"installment_status_id" db:"installment_status_id" binding:"omitempty"`
	Name                string    `json:"name" db:"name" binding:"omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
