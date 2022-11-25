package body

const (
	InvalidLoanPeriodIDFormatMessage  = "Invalid loan period id format."
	InvalidInstallmentIDFormatMessage = "Invalid installment id format."
	InvalidLoanIDFormatMessage        = "Invalid loan id format."
	InvalidAmountFormatMessage        = "Invalid amount format."
	InvalidNameFormatMessage          = "Invalid name format."
)

type UnprocessableEntity struct {
	Fields map[string]string `json:"fields"`
}
