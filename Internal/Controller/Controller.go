package Controller

import (
	"CSR/Internal/contracts"
	"CSR/Internal/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service contracts.ServiceI
	router  *gin.Engine
}

func NewController(service contracts.ServiceI) *Controller {
	return &Controller{service: service, router: gin.Default()}
}

func (ctrl *Controller) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUserNotFound) || errors.Is(err, errs.ErrNotFound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidUserID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})

	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}

}
