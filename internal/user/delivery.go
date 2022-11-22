package user

import "github.com/gin-gonic/gin"

type Handlers interface {
	DebtorDetails(c *gin.Context)
	ContractConfirm(c *gin.Context)
	CreateLoan(c *gin.Context)
}
