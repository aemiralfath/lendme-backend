package models

import (
	"github.com/google/uuid"
	"time"
)

type Installment struct {
	InstallmentID       uuid.UUID             `json:"installment_id" db:"installment_id" binding:"omitempty"`
	LendingID           uuid.UUID             `json:"lending_id" db:"lending_id" binding:"omitempty"`
	InstallmentStatusID int                   `json:"installment_status_id" db:"installment_status_id" binding:"omitempty"`
	Amount              float64               `json:"amount" db:"amount" binding:"omitempty"`
	DueDate             time.Time             `json:"due_date" db:"due_date" binding:"omitempty"`
	CreatedAt           time.Time             `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt           time.Time             `json:"updated_at,omitempty" db:"updated_at"`
	InstallmentStatus   InstallmentStatusType `json:"installment_status,omitempty" gorm:"foreignKey:InstallmentStatusID;references:InstallmentStatusID"`
}

func (i *Installment) PrepareCreate() error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	i.InstallmentID = id
	i.InstallmentStatusID = 1

	return nil
}