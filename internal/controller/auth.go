package controller

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"CSR/internal/pkg"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"net/http"
)

// SignUp
// @Summary Registration
// @Description create new account
// @Tags Auth
// @Consume json
// @Produce json
// @Param request_body body models.SignUpRequest true "Information about new account"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/sign-up [post]
func (ctrl *Controller) SignUp(c *gin.Context) {
	user := models.SignUpRequest{}
	c.BindJSON(&user)
	err := ctrl.service.CreateNewUser(user)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Info: "new user successfully created"})

}

// SignIn
// @Summary Sign in
// @Description sign in account
// @Tags Auth
// @Consume json
// @Produce json
// @Param request_body body models.SignInRequest true "sign in credentials"
// @Success 200 {object} models.TokenPairResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/sign-in [post]
func (ctrl *Controller) SignIn(c *gin.Context) {
	user := models.SignInRequest{}
	err := c.BindJSON(&user)
	fmt.Println(user)
	if err != nil {
		ctrl.handleError(c, errs.ErrInvalidRequestBody)
		return
	}
	userId,userRole,err := ctrl.service.Authenticate(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, CommonError{
			Error: err.Error(),
		})
		return
	}
	accessToken, refreshToken, err := ctrl.generateNewTokenPair(userId,userRole)
	if err != nil {
		ctrl.handleError(c, err)
	}
	c.JSON(200, models.TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

}

// RefreshTokenPair
// @Summary Refresh token pair
// @Description refresh token pair
// @Tags Auth
// @Produce json
// @Param X-Refresh-Token header string true "insert refresh token"
// @Success 200 {object} models.TokenPairResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/refresh [get]
func (ctrl *Controller) RefreshTokenPair(c *gin.Context) {
	token, err := ctrl.extractTokenFromHeader(c, refreshTokenHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}
	userID, isRefresh,role, err := pkg.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}
	log.Println("token: ", token, "isRefresh controller: ", isRefresh)
	if !isRefresh {
		c.JSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	} else {
		accessToken, refreshToken, err := ctrl.generateNewTokenPair(userID,role)
		if err != nil {
			ctrl.handleError(c, err)
		}
		c.JSON(http.StatusOK, models.TokenPairResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
	}

}
