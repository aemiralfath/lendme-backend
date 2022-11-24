package admin

import "github.com/gin-gonic/gin"

type Handlers interface {
	UpdateDebtorByID(c *gin.Context)
	GetDebtors(c *gin.Context)
	GetDebtorByID(c *gin.Context)
	GetLoans(c *gin.Context)
	GetLoanByID(c *gin.Context)
	ApproveLoan(c *gin.Context)
}
