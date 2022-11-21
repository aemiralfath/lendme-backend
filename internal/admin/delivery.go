package admin

import "github.com/gin-gonic/gin"

type Handlers interface {
	UpdateDebtorByID(c *gin.Context)
	GetDebtors(c *gin.Context)
}
