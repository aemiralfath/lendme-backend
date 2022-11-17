package models

import (
	"github.com/google/uuid"
	"time"
)

type Debtor struct {
	DebtorID           uuid.UUID            `json:"debtor_id" db:"debtor_id" binding:"omitempty"`
	UserID             uuid.UUID            `json:"user_id" db:"user_id" binding:"omitempty"`
	CreditHealthID     int                  `json:"credit_health_id" db:"user_id" binding:"omitempty"`
	ContractTrackingID int                  `json:"contract_tracking_id" db:"user_id" binding:"omitempty"`
	CreditLimit        float64              `json:"credit_limit" db:"credit_limit" binding:"omitempty"`
	CreatedAt          time.Time            `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at,omitempty" db:"updated_at"`
	CreditHealth       CreditHealthType     `json:"credit_health,omitempty" gorm:"foreignKey:CreditHealthID;references:CreditHealthID"`
	ContractTracking   ContractTrackingType `json:"contract_tracking,omitempty" gorm:"foreignKey:ContractTrackingID;references:ContractTrackingID"`
}

func (d *Debtor) PrepareCreate(userID uuid.UUID, creditHealthID, contractTrackingID int) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	d.DebtorID = id
	d.UserID = userID
	d.CreditHealthID = creditHealthID
	d.ContractTrackingID = contractTrackingID

	return nil
}