package body

const (
	InvalidLoanPeriodIDFormatMessage  = "Invalid loan period id format."
	InvalidInstallmentIDFormatMessage = "Invalid installment id format."
	InvalidLoanIDFormatMessage        = "Invalid loan id format."
	InvalidAmountFormatMessage        = "Invalid amount format."
	InvalidNameFormatMessage          = "Invalid name format."
	InvalidPhoneNumberFormatMessage   = "Invalid phone number format."
	InvalidAddressFormatMessage       = "Invalid address format."
	InvalidEmailFormatMessage         = "Invalid email format."
)

type UnprocessableEntity struct {
	Fields map[string]string `json:"fields"`
}
