package body

const (
	InvalidDebtorIDFormatMessage       = "Invalid debtor id format."
	InvalidContractStatusFormatMessage = "Invalid contract status format."
	InvalidCreditHealthFormatMessage   = "Invalid credit health format."
	InvalidCreditLimitFormatMessage    = "Invalid credit limit format."
)

type UnprocessableEntity struct {
	Fields map[string]string `json:"fields"`
}
