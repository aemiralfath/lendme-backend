package delivery

import (
	"final-project-backend/internal/admin"
	"final-project-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapAdminRoutes(adminGroup *gin.RouterGroup, h admin.Handlers, mw *middleware.MWManager) {
	adminGroup.Use(mw.AuthJWTMiddleware())
	adminGroup.Use(mw.AdminMiddleware())
	adminGroup.GET("/", h.GetSummary)
	adminGroup.GET("/debtors", h.GetDebtors)
	adminGroup.GET("/debtors/:id", h.GetDebtorByID)
	adminGroup.PUT("/debtors/:id", h.UpdateDebtorByID)
	adminGroup.GET("/loans", h.GetLoans)
	adminGroup.GET("/loans/:id", h.GetLoanByID)
	adminGroup.PUT("/loans/:id", h.ApproveLoan)
	adminGroup.DELETE("/loans/:id", h.RejectLoan)
	adminGroup.GET("/loans/installments/:id", h.GetInstallmentByID)
	adminGroup.PUT("/loans/installments/:id", h.UpdateInstallmentByID)
	adminGroup.GET("/payments", h.GetPayments)
	adminGroup.GET("/vouchers", h.GetVouchers)
	adminGroup.POST("/vouchers", h.CreateVoucher)
	adminGroup.GET("/vouchers/:id", h.GetVoucherByID)
	adminGroup.PUT("/vouchers/:id", h.UpdateVoucher)
	adminGroup.DELETE("/vouchers/:id", h.DeleteVoucher)
}
