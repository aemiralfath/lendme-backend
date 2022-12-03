package delivery

import (
	"errors"
	"final-project-backend/config"
	"final-project-backend/internal/admin"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/logger"
	"final-project-backend/pkg/response"
	"final-project-backend/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type adminHandlers struct {
	cfg     *config.Config
	adminUC admin.UseCase
	logger  logger.Logger
}

func NewAdminHandlers(cfg *config.Config, adminUC admin.UseCase, log logger.Logger) admin.Handlers {
	return &adminHandlers{cfg: cfg, adminUC: adminUC, logger: log}
}

func (h *adminHandlers) GetDebtorByID(c *gin.Context) {
	id := c.Param("id")
	debtor, err := h.adminUC.GetDebtorByID(c, id)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, debtor, http.StatusOK)
}

func (h *adminHandlers) GetVoucherByID(c *gin.Context) {
	id := c.Param("id")
	voucher, err := h.adminUC.GetVoucherByID(c, id)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, voucher, http.StatusOK)
}

func (h *adminHandlers) ApproveLoan(c *gin.Context) {
	id := c.Param("id")
	lending, err := h.adminUC.ApproveLoan(c, id)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, lending, http.StatusOK)
}

func (h *adminHandlers) RejectLoan(c *gin.Context) {
	id := c.Param("id")
	lending, err := h.adminUC.RejectLoan(c, id)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, lending, http.StatusOK)
}

func (h *adminHandlers) UpdateDebtorByID(c *gin.Context) {
	id := c.Param("id")
	var requestBody body.UpdateContractRequest
	if err := c.ShouldBind(&requestBody); err != nil {
		response.ErrorResponse(c.Writer, response.BadRequestMessage, http.StatusBadRequest)
		return
	}

	invalidFields, err := requestBody.Validate()
	if err != nil {
		response.ErrorResponseData(c.Writer, invalidFields, response.UnprocessableEntityMessage, http.StatusUnprocessableEntity)
		return
	}

	debtor, err := h.adminUC.UpdateDebtorByID(c, id, requestBody)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, debtor, http.StatusOK)
}

func (h *adminHandlers) UpdateVoucher(c *gin.Context) {
	id := c.Param("id")
	var requestBody body.UpdateVoucherRequest
	if err := c.ShouldBind(&requestBody); err != nil {
		response.ErrorResponse(c.Writer, response.BadRequestMessage, http.StatusBadRequest)
		return
	}

	invalidFields, err := requestBody.Validate()
	if err != nil {
		response.ErrorResponseData(c.Writer, invalidFields, response.UnprocessableEntityMessage, http.StatusUnprocessableEntity)
		return
	}

	voucher, err := h.adminUC.UpdateVoucherByID(c, id, requestBody)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, voucher, http.StatusOK)
}

func (h *adminHandlers) CreateVoucher(c *gin.Context) {
	var requestBody body.CreateVoucherRequest
	if err := c.ShouldBind(&requestBody); err != nil {
		response.ErrorResponse(c.Writer, response.BadRequestMessage, http.StatusBadRequest)
		return
	}

	invalidFields, err := requestBody.Validate()
	if err != nil {
		response.ErrorResponseData(c.Writer, invalidFields, response.UnprocessableEntityMessage, http.StatusUnprocessableEntity)
		return
	}

	voucher, err := h.adminUC.CreateVoucher(c, requestBody)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, voucher, http.StatusOK)
}

func (h *adminHandlers) DeleteVoucher(c *gin.Context) {
	id := c.Param("id")
	voucher, err := h.adminUC.DeleteVoucherByID(c, id)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, voucher, http.StatusOK)
}

func (h *adminHandlers) GetInstallmentByID(c *gin.Context) {
	id := c.Param("id")
	installment, err := h.adminUC.GetInstallmentByID(c, id)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, installment, http.StatusOK)
}

func (h *adminHandlers) UpdateInstallmentByID(c *gin.Context) {
	id := c.Param("id")
	var requestBody body.UpdateInstallmentRequest
	if err := c.ShouldBind(&requestBody); err != nil {
		response.ErrorResponse(c.Writer, response.BadRequestMessage, http.StatusBadRequest)
		return
	}

	invalidFields, err := requestBody.Validate()
	if err != nil {
		response.ErrorResponseData(c.Writer, invalidFields, response.UnprocessableEntityMessage, http.StatusUnprocessableEntity)
		return
	}

	installment, err := h.adminUC.UpdateInstallmentByID(c, id, requestBody)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, installment, http.StatusOK)
}

func (h *adminHandlers) GetLoanByID(c *gin.Context) {
	id := c.Param("id")
	loan, err := h.adminUC.GetLoanByID(c, id)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, loan, http.StatusOK)
}

func (h *adminHandlers) GetLoans(c *gin.Context) {
	pagination := &utils.Pagination{}
	name, status := h.ValidateQueryLoans(c, pagination)

	loans, err := h.adminUC.GetLoans(c, name, status, pagination)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, loans, http.StatusOK)
}

func (h *adminHandlers) GetSummary(c *gin.Context) {
	summary, err := h.adminUC.GetSummary(c)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, summary, http.StatusOK)
}

func (h *adminHandlers) ValidateQueryLoans(c *gin.Context, pagination *utils.Pagination) (string, []int) {
	name := strings.TrimSpace(c.Query("name"))
	status := strings.TrimSpace(c.Query("status"))
	sort := strings.TrimSpace(c.Query("sort"))
	sortBy := strings.TrimSpace(c.Query("sortBy"))
	limit := strings.TrimSpace(c.Query("limit"))
	page := strings.TrimSpace(c.Query("page"))

	var statusFilter []int
	var sortFilter string
	var sortByFilter string
	var limitFilter int
	var pageFilter int

	switch status {
	case "history":
		statusFilter = append(statusFilter, 4, 5)
	default:
		statusFilter = append(statusFilter, 1, 2, 3)
	}

	switch sort {
	case "asc":
		sortFilter = sort
	default:
		sortFilter = "desc"
	}

	switch sortBy {
	case "amount":
		sortByFilter = sortBy
	default:
		sortByFilter = "created_at"
	}

	limitFilter, err := strconv.Atoi(limit)
	if err != nil || limitFilter < 1 {
		limitFilter = 10
	}

	pageFilter, err = strconv.Atoi(page)
	if err != nil || pageFilter < 1 {
		pageFilter = 1
	}

	pagination.Limit = limitFilter
	pagination.Page = pageFilter
	pagination.Sort = fmt.Sprintf("%s %s", sortByFilter, sortFilter)

	return name, statusFilter
}

func (h *adminHandlers) GetPayments(c *gin.Context) {
	pagination := &utils.Pagination{}
	name := h.ValidateQueryPayments(c, pagination)

	payments, err := h.adminUC.GetPayments(c, name, pagination)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, payments, http.StatusOK)
}

func (h *adminHandlers) ValidateQueryPayments(c *gin.Context, pagination *utils.Pagination) string {
	name := strings.TrimSpace(c.Query("name"))
	sort := strings.TrimSpace(c.Query("sort"))
	sortBy := strings.TrimSpace(c.Query("sortBy"))
	limit := strings.TrimSpace(c.Query("limit"))
	page := strings.TrimSpace(c.Query("page"))

	var sortFilter string
	var sortByFilter string
	var limitFilter int
	var pageFilter int

	switch sort {
	case "asc":
		sortFilter = sort
	default:
		sortFilter = "desc"
	}

	switch sortBy {
	case "payment_amount":
		sortByFilter = sortBy
	default:
		sortByFilter = "payment_date"
	}

	limitFilter, err := strconv.Atoi(limit)
	if err != nil || limitFilter < 1 {
		limitFilter = 10
	}

	pageFilter, err = strconv.Atoi(page)
	if err != nil || pageFilter < 1 {
		pageFilter = 1
	}

	pagination.Limit = limitFilter
	pagination.Page = pageFilter
	pagination.Sort = fmt.Sprintf("%s %s", sortByFilter, sortFilter)

	return name
}

func (h *adminHandlers) GetVouchers(c *gin.Context) {
	pagination := &utils.Pagination{}
	name := h.ValidateQueryVouchers(c, pagination)

	vouchers, err := h.adminUC.GetVouchers(c, name, pagination)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, vouchers, http.StatusOK)
}

func (h *adminHandlers) ValidateQueryVouchers(c *gin.Context, pagination *utils.Pagination) string {
	name := strings.TrimSpace(c.Query("name"))
	sort := strings.TrimSpace(c.Query("sort"))
	sortBy := strings.TrimSpace(c.Query("sortBy"))
	limit := strings.TrimSpace(c.Query("limit"))
	page := strings.TrimSpace(c.Query("page"))

	var sortFilter string
	var sortByFilter string
	var limitFilter int
	var pageFilter int

	switch sort {
	case "asc":
		sortFilter = sort
	default:
		sortFilter = "desc"
	}

	switch sortBy {
	case "expire_date":
		sortByFilter = sortBy
	case "discount_payment":
		sortByFilter = sortBy
	case "discount_quota":
		sortByFilter = sortBy
	default:
		sortByFilter = "active_date"
	}

	limitFilter, err := strconv.Atoi(limit)
	if err != nil || limitFilter < 1 {
		limitFilter = 10
	}

	pageFilter, err = strconv.Atoi(page)
	if err != nil || pageFilter < 1 {
		pageFilter = 1
	}

	pagination.Limit = limitFilter
	pagination.Page = pageFilter
	pagination.Sort = fmt.Sprintf("%s %s", sortByFilter, sortFilter)

	return name
}

func (h *adminHandlers) GetDebtors(c *gin.Context) {
	pagination := &utils.Pagination{}
	name := h.ValidateQueryDebtors(c, pagination)

	debtors, err := h.adminUC.GetDebtors(c, name, pagination)
	if err != nil {
		var e *httperror.Error
		if !errors.As(err, &e) {
			h.logger.Errorf("HandlerRegister, Error: %s", err)
			response.ErrorResponse(c.Writer, response.InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}

		response.ErrorResponse(c.Writer, e.Err.Error(), e.Status)
		return
	}

	response.SuccessResponse(c.Writer, debtors, http.StatusOK)
}

func (h *adminHandlers) ValidateQueryDebtors(c *gin.Context, pagination *utils.Pagination) string {
	name := strings.TrimSpace(c.Query("name"))
	sort := strings.TrimSpace(c.Query("sort"))
	sortBy := strings.TrimSpace(c.Query("sortBy"))
	limit := strings.TrimSpace(c.Query("limit"))
	page := strings.TrimSpace(c.Query("page"))

	var sortFilter string
	var sortByFilter string
	var limitFilter int
	var pageFilter int

	switch sort {
	case "asc":
		sortFilter = sort
	default:
		sortFilter = "desc"
	}

	switch sortBy {
	case "credit_limit":
		sortByFilter = sortBy
	case "credit_used":
		sortByFilter = sortBy
	case "total_delay":
		sortByFilter = sortBy
	default:
		sortByFilter = "created_at"
	}

	limitFilter, err := strconv.Atoi(limit)
	if err != nil || limitFilter < 1 {
		limitFilter = 10
	}

	pageFilter, err = strconv.Atoi(page)
	if err != nil || pageFilter < 1 {
		pageFilter = 1
	}

	pagination.Limit = limitFilter
	pagination.Page = pageFilter
	pagination.Sort = fmt.Sprintf("%s %s", sortByFilter, sortFilter)

	return name
}
