package middleware

import (
	"final-project-backend/pkg/response"
	"final-project-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (mw *MWManager) AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, err := utils.ExtractJWTFromRequest(c.Request, mw.cfg.Server.JwtSecretKey)
		if err != nil {
			response.ErrorResponse(c.Writer, response.ForbiddenMessage, http.StatusForbidden)
			c.Abort()
			return
		}

		mw.logger.Infof("body middleware bearerHeader %s", claim["id"].(string))
		c.Set("userID", claim["id"].(string))
		c.Set("roleID", claim["role_id"].(float64))
		c.Next()
	}
}
