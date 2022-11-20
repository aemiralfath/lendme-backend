package delivery

import (
	"errors"
	"final-project-backend/config"
	"final-project-backend/internal/user"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/logger"
	"final-project-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandlers struct {
	cfg    *config.Config
	userUC user.UseCase
	logger logger.Logger
}

func NewUserHandlers(cfg *config.Config, userUC user.UseCase, log logger.Logger) user.Handlers {
	return &userHandlers{cfg: cfg, userUC: userUC, logger: log}
}

func (h *userHandlers) ContractConfirm(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		response.ErrorResponse(c.Writer, response.UnauthorizedMessage, http.StatusUnauthorized)
		return
	}

	debtor, err := h.userUC.ConfirmContract(c, userID.(string))
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

func (h *userHandlers) DebtorDetails(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		response.ErrorResponse(c.Writer, response.UnauthorizedMessage, http.StatusUnauthorized)
		return
	}

	userDebtor, err := h.userUC.GetDebtorDetails(c, userID.(string))
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

	response.SuccessResponse(c.Writer, userDebtor, http.StatusOK)
}
