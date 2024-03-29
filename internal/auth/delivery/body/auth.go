package body

const (
	InvalidNameFormatMessage        = "Invalid name format."
	InvalidPhoneNumberFormatMessage = "Invalid phone number format."
	InvalidAddressFormatMessage     = "Invalid address format."
	InvalidEmailFormatMessage       = "Invalid email format."
	InvalidPasswordFormatMessage    = "Password must contain at least 8-40 characters," +
		"at least 1 number, 1 Upper case, and 1 special character"
)

type UnprocessableEntity struct {
	Fields map[string]string `json:"fields"`
}
