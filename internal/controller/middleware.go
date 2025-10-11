package controller

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"CSR/internal/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
const(
	authorizationHeader="Authorization"
	userIDCtx="User ID"
	refreshTokenHeader="X-Refresh-Token"
	userRoleCtx="userRole"

)

func (ctrl*Controller)checkUserAuthentication(c*gin.Context){
	token,err:= ctrl.extractTokenFromHeader(c,authorizationHeader)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,CommonError{Error: err.Error()})
		return 
	}
	userID,isRefresh,role,err:=pkg.ParseToken(token)
	log.Println(userID,isRefresh,role)
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
	c.Set(userRoleCtx,string(role))
	
}
func (ctrl*Controller)CheckIsAdmin(c*gin.Context){
	role:=c.GetString(userRoleCtx)
	if role==""{
		c.AbortWithStatusJSON(http.StatusUnauthorized,CommonError{Error: "role is null "})
		return 
	}
	if role!=models.RoleAdmin{
		c.AbortWithStatusJSON(http.StatusForbidden,CommonError{Error: errs.ErrAccessDenied.Error()})
		return 
	}
	c.Next()
	
}