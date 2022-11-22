package body

import (
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"net/http"
)

type CreateLoan struct {
	LoadPeriodID int     `json:"loan_period_id"`
	Amount       float64 `json:"amount"`
}

func (r *CreateLoan) Validate() (UnprocessableEntity, error) {
	unprocessableEntity := false
	entity := UnprocessableEntity{
		Fields: map[string]string{
			"loan_period_id": "",
			"amount":         "",
		},
	}

	if r.LoadPeriodID == 0 {
		unprocessableEntity = true
		entity.Fields["loan_period_id"] = InvalidLoanPeriodIDFormatMessage
	}

	if r.Amount < 1000000 {
		unprocessableEntity = true
		entity.Fields["amount"] = InvalidAmountFormatMessage
	}

	if unprocessableEntity {
		return entity, httperror.New(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityMessage,
		)
	}

	return entity, nil
}
