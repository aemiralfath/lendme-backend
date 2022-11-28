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
	adminGroup.GET("/debtors/:id", h.GetDebtorByID)
	adminGroup.PUT("/debtors/:id", h.UpdateDebtorByID)
	adminGroup.GET("/loans", h.GetLoans)
	adminGroup.GET("/loans/:id", h.GetLoanByID)
	adminGroup.PUT("/loans/:id", h.ApproveLoan)
	adminGroup.GET("/loans/installment/:id", h.GetInstallmentByID)
	adminGroup.PUT("/loans/installment/:id", h.UpdateInstallmentByID)
	adminGroup.GET("/payments", h.GetPayments)
	adminGroup.POST("/vouchers", h.CreateVoucher)
}
