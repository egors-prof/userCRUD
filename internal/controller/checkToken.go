package controller

import (
	"CSR/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type checkToken struct{
	TokenString string `json:"token_string"`
}

func (ctrl Controller) checkToken(c *gin.Context) {
	var token checkToken=checkToken{}
	err:=c.BindJSON(&token)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,CommonError{Error:err.Error()})
		return 
	}
	id,isRefresh,err:=pkg.ParseToken(token.TokenString)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,CommonError{Error:err.Error()})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"id":id,
		"is_refresh":isRefresh,
	})



}
	
