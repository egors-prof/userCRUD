package controller

import (
	"CSR/internal/pkg"
	"log"
	"net/http"


	"github.com/gin-gonic/gin"
)
const(
	authorizationHeader="Authorization"
	userIDCtx="User ID"
	refreshTokenHeader="X-Refresh-Token"

)

func (ctrl*Controller)checkUserAuthentication(c*gin.Context){
	token,err:= ctrl.extractTokenFromHeader(c,authorizationHeader)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,CommonError{Error: err.Error()})
		return 
	}
	userID,isRefresh,err:=pkg.ParseToken(token)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized,CommonError{Error: err.Error()})
		return 
	}

	log.Println("token: ",token,"isRefresh controller: ",isRefresh)
	if isRefresh{
		c.AbortWithStatusJSON(http.StatusUnauthorized,CommonError{Error: "inappropriate token"})
		return 
	}
	c.Set(userIDCtx,userID)
}