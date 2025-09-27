package controller

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateUserRequest struct {
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Age   int    `db:"age" json:"age"`
}

// GetAllUsers
// @Summary getting users
// @Description getting list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users [get]
func (ctrl *Controller) GetAllUsers(c *gin.Context) {

	users, err := ctrl.service.GetAllUsers()
	if err != nil {
		c.JSON(400, CommonError{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

// GetUserById
// @Summary getting user
// @Description getting a user by id
// @Tags Users
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} models.User
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [get]
func (ctrl *Controller) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, err)
		return
	}
	user, err := ctrl.service.GetUserById(id)
	if err != nil {
		ctrl.handleError(c, errs.ErrUserNotFound)
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
func (ctrl *Controller) CreateNewUser(c *gin.Context) {
	var newUser models.User
	err := c.BindJSON(&newUser)
	if err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidUserID, err))
		return
	}
	if newUser.Age == 0 || newUser.Name == "" || newUser.Email == "" {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
	}
	err = ctrl.service.CreateNewUser(newUser)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, DefaultResponse{Info: "new user created"})

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
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		ctrl.handleError(c, err)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctrl.handleError(c, err)
		return
	}
	err = ctrl.service.UpdateUserById(id, user)
	if err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidUserID, err))
		return
	}
	c.JSON(http.StatusOK, DefaultResponse{Info: "user info is successfully updated"})

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
	if err != nil || id < 1 {
		ctrl.handleError(c, err)
		return
	}
	err = ctrl.service.DeleteUserById(id)
	if err != nil {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return

	}
	c.JSON(http.StatusOK, DefaultResponse{Info: "user is successfully deleted"})

}
