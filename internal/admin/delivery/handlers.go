package delivery

import (
	"errors"
	"final-project-backend/config"
	"final-project-backend/internal/admin"
	"final-project-backend/internal/admin/delivery/body"
	"final-project-backend/pkg/httperror"
	"final-project-backend/pkg/logger"
	"final-project-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type adminHandlers struct {
	cfg     *config.Config
	adminUC admin.UseCase
	logger  logger.Logger
}

func NewAdminHandlers(cfg *config.Config, adminUC admin.UseCase, log logger.Logger) admin.Handlers {
	return &adminHandlers{cfg: cfg, adminUC: adminUC, logger: log}
}

func (h *adminHandlers) UpdateContractStatus(c *gin.Context) {
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

	debtor, err := h.adminUC.UpdateContractStatus(c, requestBody)
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
