package response

import (
	"encoding/json"
	"net/http"
)

const (
	VersionMessage             = "version 1"
	BadRequestMessage          = "Invalid request."
	UnprocessableEntityMessage = "Unprocessable entity."
	InternalServerErrorMessage = "Something is wrong, pls try again later."
	NotFoundMessage            = "Route does not exist, please check again your route path."
	UnauthorizedMessage        = "Unauthorized"
	ForbiddenMessage           = "Forbidden"

	EmailAlreadyExistMessage           = "Email already exist."
	DebtorIDNotExist                   = "Debtor ID not exist."
	UserIDNotExist                     = "User ID not exist."
	LendingIDNotExist                  = "Lending ID not exist."
	ContractIDNotExist                 = "Contract ID not exist."
	ContractNotAccepted                = "Contract not accepted."
	ContractNotConfirmed               = "Contract not confirmed."
	ContractAlreadyAccepted            = "Contract already accepted."
	LoanPeriodNotExist                 = "Loan period ID not exist."
	InstallmentNotExist                = "Installment ID not exist."
	VoucherNotExist                    = "Voucher ID not exist."
	LendingInstallmentNotMatch         = "Lending installment not match"
	InstallmentAlreadyPaid             = "Installment already paid."
	LoanAmountExceedCreditLimit        = "Loan amount exceed credit limit."
	LoanAmountExceedCreditLimitWarning = "Loan amount exceed credit limit warning."
	CreditHealthStatusBlocked          = "Credit health status blocked"
)

type JSONResponse struct {
	Message string      `response:"message,omitempty"`
	Data    interface{} `response:"data,omitempty"`
}

func returnJSONResponse(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(JSONResponse{
		Message: message,
		Data:    data,
	})
}

func SuccessResponse(w http.ResponseWriter, data interface{}, statusCode ...int) {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	returnJSONResponse(w, "success", data, code)
}

func ErrorResponse(w http.ResponseWriter, message string, statusCode ...int) {
	code := http.StatusBadRequest
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	returnJSONResponse(w, message, nil, code)
}

func ErrorResponseData(w http.ResponseWriter, data interface{}, message string, statusCode ...int) {
	code := http.StatusBadRequest
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	returnJSONResponse(w, message, data, code)
}
