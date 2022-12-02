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
	userGroup.PUT("/details", h.UpdateUser)
	userGroup.PATCH("/details", h.ContractConfirm)
	userGroup.GET("/loans", h.GetLoans)
	userGroup.POST("/loans", h.CreateLoan)
	userGroup.GET("/loans/:id", h.GetLoanByID)
	userGroup.GET("/loans/installments/:id", h.GetInstallmentByID)
	userGroup.POST("/loans/installments/:id", h.CreatePayment)
	userGroup.GET("/vouchers", h.GetVouchers)
	userGroup.GET("/payments", h.GetPayments)
}
