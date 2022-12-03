package middleware

import (
	"final-project-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (mw *MWManager) AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, exist := c.Get("roleID")
		if !exist || roleID.(float64) != 1 {
			response.ErrorResponse(c.Writer, response.ForbiddenMessage, http.StatusForbidden)
			c.Abort()
			return
		}
		c.Next()
	}
}
