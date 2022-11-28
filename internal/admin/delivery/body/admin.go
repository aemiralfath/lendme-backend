package body

const (
	InvalidContractStatusFormatMessage = "Invalid contract status format."
	InvalidCreditHealthFormatMessage   = "Invalid credit health format."
	InvalidCreditLimitFormatMessage    = "Invalid credit limit format."
	InvalidDateFormatMessage           = "Invalid due date format."
	InvalidNameFormatMessage           = "Invalid name format."
	InvalidDiscountFormatMessage       = "Invalid discount format."
)

type UnprocessableEntity struct {
	Fields map[string]string `json:"fields"`
}
