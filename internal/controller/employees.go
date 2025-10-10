package controller

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"errors"
	"log"
	"os"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// GetAllEmployees
// @Summary getting employees
// @Description getting list of all employees
// @Tags Employees
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Employee
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/employees [get]
func (ctrl *Controller) GetAllEmployees(c *gin.Context) {
	logger := zerolog.New(os.Stdout).With().Str("func_name", "controller.GetAllEmployees").Logger()
	userID := c.GetInt(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid user id in context"})
	}
	logger.Debug().Int("user_id", userID).Msg("GetUser")
	emps, err := ctrl.service.GetAllEmployees()
	if err != nil {
		ctrl.handleError(c, err)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	c.JSON(http.StatusOK, emps)

}

// GetEmployeeById
// @Summary getting employee
// @Description getting an employee by id
// @Tags Employees
// @Produce json
// @Security BearerAuth
// @Param id path int true "employee id"
// @Success 200 {object} models.Employee
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/employees/{id} [get]
func (ctrl *Controller) GetEmployeeById(c *gin.Context) {
	logger := zerolog.New(os.Stdout).With().Str("func_name", "controller.GetEmployeeById").Logger()
	userID := c.GetInt(userIDCtx)
	log.Println("userID context: ", userID)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid user id in context"})
	}
	logger.Debug().Int("user_id", userID).Msg("GetUser")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	user, err := ctrl.service.GetEmployeeById(id)
	if err != nil {
		ctrl.handleError(c, errs.ErrUserNotFound)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	c.JSON(200, user)
}

// CreateNewEmployee
// @Summary creating
// @Description creating new employee
// @Tags Employees
// @Produce json
// @Consume json
// @Security BearerAuth
// @Param request_body body models.EmployeeRequest true "user info"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/employees [post]
func (ctrl *Controller) CreateNewEmployee(c *gin.Context) {
	var newEmp models.EmployeeRequest
	err := c.BindJSON(&newEmp)
	if err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidUserID, err))
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	if newEmp.Age == 0 || newEmp.Name == "" || newEmp.Email == "" {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	err = ctrl.service.CreateNewEmployee(newEmp)
	if err != nil {
		ctrl.handleError(c, err)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	c.JSON(http.StatusOK, CommonResponse{Info: "new employee created"})

}

// UpdateUserById
// @Summary updating
// @Description updating user by id
// @Tags Employees
// @Produce json
// @Consume json
// @Security BearerAuth
// @Param id path int true "user id"
// @Param request_body body models.EmployeeRequest true "user info"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/employees/{id} [put]
func (ctrl *Controller) UpdateUserById(c *gin.Context) {
	var empRequest models.EmployeeRequest
	if err := c.BindJSON(&empRequest); err != nil {
		ctrl.handleError(c, err)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}

	if empRequest.Age == 0 && empRequest.Email == "" && empRequest.Name == "" {
		ctrl.handleError(c, errs.ErrInvalidRequestBody)
		ctrl.ctrlLogger.Error().Err(errs.ErrInvalidRequestBody).Send()
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, err)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	err = ctrl.service.UpdateEmployeeById(id, empRequest)
	if err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidUserID, err))
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	c.JSON(http.StatusOK, CommonResponse{Info: "user info is successfully updated"})

}

// DeleteEmployeeById
// @Summary deleting emp
// @Description delete employee by id
// @Tags Employees
// @Produce json
// @Security BearerAuth
// @Param id path int true "emp id"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/employees/{id} [delete]
func (ctrl *Controller) DeleteEmployeeById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.handleError(c, errs.ErrInvalidIDFormat)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	if id < 1 {
		ctrl.handleError(c, errs.ErrNegativeID)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}

	err = ctrl.service.DeleteEmployeeById(id)
	if err != nil {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return

	}
	c.JSON(http.StatusOK, CommonResponse{Info: "employee is successfully deleted"})

}
