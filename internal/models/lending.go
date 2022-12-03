package models

import (
	"github.com/google/uuid"
	"time"
)

type Lending struct {
	LendingID       uuid.UUID          `json:"lending_id" db:"lending_id" binding:"omitempty"`
	DebtorID        uuid.UUID          `json:"debtor_id" db:"debtor_id" binding:"omitempty"`
	LoanPeriodID    int                `json:"loan_period_id" db:"loan_period_id" binding:"omitempty"`
	LendingStatusID int                `json:"lending_status_id" db:"lending_status_id" binding:"omitempty"`
	Name            string             `json:"name" db:"name" binding:"omitempty"`
	Amount          float64            `json:"amount" db:"amount" binding:"omitempty"`
	CreatedAt       time.Time          `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at,omitempty" db:"updated_at"`
	Debtor          *Debtor            `json:"debtor,omitempty" gorm:"foreignKey:DebtorID;references:DebtorID"`
	LoanPeriod      *LoanPeriod        `json:"loan_period,omitempty" gorm:"foreignKey:LoanPeriodID;references:LoanPeriodID"`
	LendingStatus   *LendingStatusType `json:"lending_status,omitempty" gorm:"foreignKey:LendingStatusID;references:LendingStatusID"`
	Installments    *[]Installment     `json:"installments,omitempty" gorm:"foreignKey:LendingID;references:LendingID"`
}

func (l *Lending) PrepareCreate() error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	l.LendingID = id
	l.LendingStatusID = 1

	return nil
}
