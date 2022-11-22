package body

import (
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"net/http"
	"strings"
)

type ApproveLoanRequest struct {
	LendingID string `json:"lending_id"`
}

func (r *ApproveLoanRequest) Validate() (UnprocessableEntity, error) {
	unprocessableEntity := false
	entity := UnprocessableEntity{
		Fields: map[string]string{
			"lending_id": "",
		},
	}

	r.LendingID = strings.TrimSpace(r.LendingID)
	if r.LendingID == "" {
		unprocessableEntity = true
		entity.Fields["lending_id"] = InvalidLendingIDFormatMessage
	}

	if unprocessableEntity {
		return entity, httperror.New(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityMessage,
		)
	}

	return entity, nil
}
