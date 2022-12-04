package body

import (
	"final-project-backend/internal/utils"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/response"
	"net/http"
	"net/mail"
	"strings"
)

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (r *RegisterRequest) Validate() (UnprocessableEntity, error) {
	unprocessableEntity := false
	entity := UnprocessableEntity{
		Fields: map[string]string{
			"name":     "",
			"email":    "",
			"password": "",
		},
	}

	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		unprocessableEntity = true
		entity.Fields["name"] = InvalidNameFormatMessage
	}

	r.PhoneNumber = strings.TrimSpace(r.PhoneNumber)
	if r.PhoneNumber == "" {
		unprocessableEntity = true
		entity.Fields["phone_number"] = InvalidPhoneNumberFormatMessage
	}

	r.Address = strings.TrimSpace(r.Address)
	if r.Address == "" {
		unprocessableEntity = true
		entity.Fields["address"] = InvalidAddressFormatMessage
	}

	r.Email = strings.TrimSpace(r.Email)
	if r.Email == "" {
		unprocessableEntity = true
		entity.Fields["email"] = InvalidEmailFormatMessage
	}

	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		unprocessableEntity = true
		entity.Fields["email"] = InvalidEmailFormatMessage
	}

	if !utils.VerifyPassword(r.Password) {
		unprocessableEntity = true
		entity.Fields["password"] = InvalidPasswordFormatMessage
	}

	if unprocessableEntity {
		return entity, httperror.New(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityMessage,
		)
	}

	return entity, nil
}
