package body

import (
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"net/http"
	"strings"
)

type CreatePayment struct {
	LendingID     string `json:"lending_id"`
	InstallmentID string `json:"installment_id"`
	VoucherID     string `json:"voucher_id"`
}

func (r *CreatePayment) Validate() (UnprocessableEntity, error) {
	unprocessableEntity := false
	entity := UnprocessableEntity{
		Fields: map[string]string{
			"lending_id":     "",
			"installment_id": "",
			"voucher_id":     "",
		},
	}

	r.VoucherID = strings.TrimSpace(r.VoucherID)

	r.LendingID = strings.TrimSpace(r.LendingID)
	if r.LendingID == "" {
		unprocessableEntity = true
		entity.Fields["lending_id"] = InvalidLoanIDFormatMessage
	}

	r.InstallmentID = strings.TrimSpace(r.InstallmentID)
	if r.InstallmentID == "" {
		unprocessableEntity = true
		entity.Fields["installment_id"] = InvalidInstallmentIDFormatMessage
	}

	if unprocessableEntity {
		return entity, httperror.New(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityMessage,
		)
	}

	return entity, nil
}
