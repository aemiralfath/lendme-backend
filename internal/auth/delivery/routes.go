package delivery

import (
	"final-project-backend/internal/auth"
	"final-project-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(authGroup *gin.RouterGroup, h auth.Handlers, mw *middleware.MWManager) {
	authGroup.POST("/register", h.Register)
	authGroup.POST("/login", h.Login)

	authGroup.Use(mw.AuthJWTMiddleware())
	authGroup.GET("/details", h.UserDetails)
}
