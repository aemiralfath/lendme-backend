package admin

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetDebtors(c *gin.Context)
	GetDebtorByID(c *gin.Context)
	GetLoans(c *gin.Context)
	GetLoanByID(c *gin.Context)
	ApproveLoan(c *gin.Context)
	RejectLoan(c *gin.Context)
	GetPayments(c *gin.Context)
	GetInstallmentByID(c *gin.Context)
	UpdateInstallmentByID(c *gin.Context)
	UpdateDebtorByID(c *gin.Context)
	CreateVoucher(c *gin.Context)
	GetVouchers(c *gin.Context)
	GetVoucherByID(c *gin.Context)
	DeleteVoucher(c *gin.Context)
	UpdateVoucher(c *gin.Context)
	GetSummary(c *gin.Context)
}
