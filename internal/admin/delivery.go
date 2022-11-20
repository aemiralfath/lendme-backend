package admin

import "github.com/gin-gonic/gin"

type Handlers interface {
	UpdateContractStatus(c *gin.Context)
}
