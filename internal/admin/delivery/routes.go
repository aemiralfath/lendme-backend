package delivery

import (
	"final-project-backend/internal/admin"
	"final-project-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapAdminRoutes(adminGroup *gin.RouterGroup, h admin.Handlers, mw *middleware.MWManager) {
	adminGroup.Use(mw.AuthJWTMiddleware())
	adminGroup.Use(mw.AdminMiddleware())
	adminGroup.GET("/debtors", h.GetDebtors)
	adminGroup.PUT("/debtors", h.UpdateDebtorByID)
}
