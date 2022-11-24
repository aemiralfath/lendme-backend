package delivery

import (
	"final-project-backend/internal/middleware"
	"final-project-backend/internal/user"
	"github.com/gin-gonic/gin"
)

func MapUserRoutes(userGroup *gin.RouterGroup, h user.Handlers, mw *middleware.MWManager) {
	userGroup.Use(mw.AuthJWTMiddleware())
	userGroup.Use(mw.UserMiddleware())
	userGroup.GET("/details", h.DebtorDetails)
	userGroup.PUT("/contract", h.ContractConfirm)
	userGroup.POST("/loans", h.CreateLoan)
	userGroup.GET("/loans", h.GetLoans)
	userGroup.GET("/loans/:id", h.GetLoanByID)
}
