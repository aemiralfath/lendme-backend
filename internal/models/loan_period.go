package models

import "time"

type LoanPeriod struct {
	LoanPeriodID int       `json:"loan_period_id" db:"loan_period_id" binding:"omitempty"`
	Duration     int       `json:"duration" db:"duration" binding:"omitempty"`
	Percentage   int       `json:"percentage" db:"percentage" binding:"omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
