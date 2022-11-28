package body

import (
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"net/http"
	"strings"
	"time"
)

type CreateVoucherRequest struct {
	Name            string    `json:"name"`
	DiscountPayment int       `json:"discount_payment"`
	DiscountQuota   int       `json:"discount_quota"`
	ActiveDate      string    `json:"active_date"`
	ExpireDate      string    `json:"expire_date"`
	ActiveDateTime  time.Time `json:"-"`
	ExpireDateTime  time.Time `json:"-"`
}

func (r *CreateVoucherRequest) Validate() (UnprocessableEntity, error) {
	unprocessableEntity := false
	entity := UnprocessableEntity{
		Fields: map[string]string{
			"name":             "",
			"discount_payment": "",
			"discount_quota":   "",
			"active_date":      "",
			"expire_date":      "",
		},
	}

	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		unprocessableEntity = true
		entity.Fields["name"] = InvalidNameFormatMessage
	}

	if r.DiscountPayment < 1 || r.DiscountPayment > 100 {
		unprocessableEntity = true
		entity.Fields["discount_payment"] = InvalidDiscountFormatMessage
	}

	if r.DiscountQuota == 0 {
		unprocessableEntity = true
		entity.Fields["discount_quota"] = InvalidDiscountFormatMessage
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")

	r.ActiveDate = strings.TrimSpace(r.ActiveDate)
	activeTime, err := time.ParseInLocation("02-01-2006 15:04:05", r.ActiveDate, loc)
	if err != nil {
		unprocessableEntity = true
		entity.Fields["active_date"] = InvalidDateFormatMessage
	}

	r.ExpireDate = strings.TrimSpace(r.ExpireDate)
	expireTime, err := time.ParseInLocation("02-01-2006 15:04:05", r.ExpireDate, loc)
	if err != nil {
		unprocessableEntity = true
		entity.Fields["expire_date"] = InvalidDateFormatMessage
	}

	if expireTime.Sub(activeTime) < 0 {
		unprocessableEntity = true
		entity.Fields["active_date"] = InvalidDateFormatMessage
		entity.Fields["expire_date"] = InvalidDateFormatMessage
	}

	r.ActiveDateTime = activeTime
	r.ExpireDateTime = expireTime
	if unprocessableEntity {
		return entity, httperror.New(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityMessage,
		)
	}

	return entity, nil
}
