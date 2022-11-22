package body

const (
	InvalidLoanPeriodIDFormatMessage = "Invalid loan period id format."
	InvalidAmountFormatMessage       = "Invalid amount format."
)

type UnprocessableEntity struct {
	Fields map[string]string `json:"fields"`
}
