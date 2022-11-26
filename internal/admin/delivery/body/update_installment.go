package body

import (
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"net/http"
	"strings"
	"time"
)

type UpdateInstallmentRequest struct {
	InstallmentID string    `json:"installment_id"`
	DueDate       string    `json:"due_date"`
	DueDateTime   time.Time `json:"-"`
}

func (r *UpdateInstallmentRequest) Validate() (UnprocessableEntity, error) {
	unprocessableEntity := false
	entity := UnprocessableEntity{
		Fields: map[string]string{
			"installment_id": "",
			"due_date":       "",
		},
	}

	r.InstallmentID = strings.TrimSpace(r.InstallmentID)
	if r.InstallmentID == "" {
		unprocessableEntity = true
		entity.Fields["installment_id"] = InvalidInstallmentIDFormatMessage
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	r.DueDate = strings.TrimSpace(r.DueDate)
	t, err := time.ParseInLocation("02-01-2006 15:04:05", r.DueDate, loc)
	if err != nil {
		unprocessableEntity = true
		entity.Fields["due_date"] = InvalidDueDateFormatMessage
	}

	r.DueDateTime = t
	if unprocessableEntity {
		return entity, httperror.New(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityMessage,
		)
	}

	return entity, nil
}
