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

// GetAllUsers
// @Summary getting employees
// @Description getting list of all employees
// @Tags Employees
// @Produce json
// @Success 200 {array} models.Employee
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/employees [get]
func (ctrl *Controller) GetAllEmployees(c *gin.Context) {
	logger:=zerolog.New(os.Stdout).With().Str("func_name","controller.GetAllEmployees").Logger()
	userID:=c.GetInt(userIDCtx)
	if userID==0{
		c.JSON(http.StatusBadRequest,CommonError{Error: "invalid user id in context"})
	}
	logger.Debug().Int("user_id",userID).Msg("GetUser")
	emps, err := ctrl.service.GetAllEmployees()
	if err != nil {
		ctrl.handleError(c,err)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	c.JSON(http.StatusOK, emps)

}

// GetUserById
// @Summary getting user
// @Description getting a user by id
// @Tags Users
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} models.Employee
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [get]
func (ctrl *Controller) GetEmployeeById(c *gin.Context) {
	logger:=zerolog.New(os.Stdout).With().Str("func_name","controller.GetEmployeeById").Logger()
	userID:=c.GetInt(userIDCtx)
	log.Println("userID context: ",userID)
	if userID==0{
		c.JSON(http.StatusBadRequest,CommonError{Error: "invalid user id in context"})
	}
	logger.Debug().Int("user_id",userID).Msg("GetUser")
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

// CreateNewUser
// @Summary creating
// @Description creating new user
// @Tags Users
// @Produce json
// @Consume json
// @Param request_body body CreateUserRequest true "user info"
// @Success 201 {object} DefaultResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users [post]
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
// @Tags Users
// @Produce json
// @Consume json
// @Param id path int true "user id"
// @Param request_body body CreateUserRequest true "user info"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [put]
func (ctrl *Controller) UpdateUserById(c *gin.Context) {
	var empRequest models.EmployeeRequest
	if err := c.BindJSON(&empRequest); err != nil {
		ctrl.handleError(c, err)
		ctrl.ctrlLogger.Error().Err(err).Send()
		return
	}
	
	if empRequest.Age==0&&empRequest.Email==""&&empRequest.Name==""{
		ctrl.handleError(c,errs.ErrInvalidRequestBody)
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

// DeleteUserById
// @Summary deleting
// @Description delete user by id
// @Tags Users
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [delete]
func (ctrl *Controller) DeleteUserById(c *gin.Context) {
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
	c.JSON(http.StatusOK, CommonResponse{Info: "user is successfully deleted"})

}
