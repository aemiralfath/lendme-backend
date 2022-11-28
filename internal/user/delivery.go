package user

import "github.com/gin-gonic/gin"

type Handlers interface {
	DebtorDetails(c *gin.Context)
	ContractConfirm(c *gin.Context)
	CreateLoan(c *gin.Context)
	GetLoans(c *gin.Context)
	GetLoanByID(c *gin.Context)
	CreatePayment(c *gin.Context)
	GetInstallmentByID(c *gin.Context)
	GetVouchers(c *gin.Context)
}
