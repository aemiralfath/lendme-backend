package middleware

import (
	"final-project-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (mw *MWManager) UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, exist := c.Get("roleID")
		if !exist || roleID.(float64) != 2 {
			response.ErrorResponse(c.Writer, response.UnauthorizedMessage, http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
