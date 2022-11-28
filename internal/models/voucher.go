package models

import (
	"github.com/google/uuid"
	"time"
)

type Voucher struct {
	VoucherID       uuid.UUID `json:"voucher_id" db:"voucher_id" binding:"omitempty"`
	Name            string    `json:"name" db:"name" binding:"omitempty"`
	DiscountPayment int       `json:"discount_payment" db:"discount_payment" binding:"omitempty"`
	DiscountQuota   int       `json:"discount_quota" db:"discount_quota" binding:"omitempty"`
	ActiveDate      time.Time `json:"active_date,omitempty" db:"active_date"`
	ExpireDate      time.Time `json:"expire_date,omitempty" db:"expire_date"`
	CreatedAt       time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (v *Voucher) PrepareCreate() error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	v.VoucherID = id

	return nil
}