package controller

import (
	"CSR/internal/errs"
	"CSR/internal/models"

	

	"github.com/gin-gonic/gin"
)
func(ctrl *Controller)CreateNewUser(c*gin.Context){
	userRequest:=models.SignUpRequest{}
	err:=c.BindJSON(&userRequest)
	if err!=nil{
		ctrl.handleError(c,errs.ErrBindJson)
		return 
	}
	
}

