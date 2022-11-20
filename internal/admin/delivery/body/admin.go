package body

const (
	InvalidDebtorIDFormatMessage       = "Invalid debtor id format."
	InvalidContractStatusFormatMessage = "Invalid contract status format."
)

type UnprocessableEntity struct {
	Fields map[string]string `json:"fields"`
}
