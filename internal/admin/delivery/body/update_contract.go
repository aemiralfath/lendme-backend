package body

import (
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"net/http"
	"strings"
)

type UpdateContractRequest struct {
	DebtorID         string  `json:"debtor_id"`
	CreditLimit      float64 `json:"credit_limit"`
	CreditHealthID   int     `json:"credit_health_id"`
	ContractStatusID int     `json:"contract_status_id"`
}

func (r *UpdateContractRequest) Validate() (UnprocessableEntity, error) {
	unprocessableEntity := false
	entity := UnprocessableEntity{
		Fields: map[string]string{
			"debtor_id":          "",
			"credit_limit":       "",
			"credit_health_id":   "",
			"contract_status_id": "",
		},
	}

	r.DebtorID = strings.TrimSpace(r.DebtorID)
	if r.DebtorID == "" {
		unprocessableEntity = true
		entity.Fields["debtor_id"] = InvalidDebtorIDFormatMessage
	}

	if r.CreditLimit < 0 {
		unprocessableEntity = true
		entity.Fields["credit_limit"] = InvalidCreditLimitFormatMessage
	}

	if r.CreditHealthID == 0 {
		unprocessableEntity = true
		entity.Fields["credit_health_id"] = InvalidCreditHealthFormatMessage
	}

	if r.ContractStatusID == 0 {
		unprocessableEntity = true
		entity.Fields["contract_status_id"] = InvalidContractStatusFormatMessage
	}

	if unprocessableEntity {
		return entity, httperror.New(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityMessage,
		)
	}

	return entity, nil
}
