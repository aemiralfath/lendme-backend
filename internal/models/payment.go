package models

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	PaymentID       uuid.UUID    `json:"payment_id" db:"payment_id" binding:"omitempty"`
	InstallmentID   uuid.UUID    `json:"installment_id" db:"installment_id" binding:"omitempty"`
	VoucherID       *uuid.UUID   `json:"voucher_id" db:"voucher_id" binding:"omitempty"`
	PaymentFine     float64      `json:"payment_fine" db:"payment_fine" binding:"omitempty"`
	PaymentDiscount float64      `json:"payment_discount" db:"payment_discount" binding:"omitempty"`
	PaymentAmount   float64      `json:"payment_amount" db:"payment_amount" binding:"omitempty"`
	PaymentDate     *time.Time   `json:"payment_date" db:"payment_date" binding:"omitempty"`
	Installment     *Installment `json:"installment,omitempty" gorm:"foreignKey:InstallmentID;references:InstallmentID"`
	Voucher         *Voucher     `json:"voucher,omitempty" gorm:"foreignKey:VoucherID;references:VoucherID"`
}

func (p *Payment) PrepareCreate() error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	p.PaymentID = id
	return nil
}
