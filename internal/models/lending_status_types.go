package models

import "time"

type LendingStatusType struct {
	LendingStatusID int       `json:"lending_status_id" db:"lending_status_id" binding:"omitempty"`
	Name            string    `json:"name" db:"name" binding:"omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
