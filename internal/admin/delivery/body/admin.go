package body

const (
	InvalidDebtorIDFormatMessage       = "Invalid debtor id format."
	InvalidInstallmentIDFormatMessage  = "Invalid installment id format."
	InvalidContractStatusFormatMessage = "Invalid contract status format."
	InvalidCreditHealthFormatMessage   = "Invalid credit health format."
	InvalidCreditLimitFormatMessage    = "Invalid credit limit format."
	InvalidDueDateFormatMessage        = "Invalid due date format."
)

type UnprocessableEntity struct {
	Fields map[string]string `json:"fields"`
}
