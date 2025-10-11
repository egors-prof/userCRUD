package controller

import (
	"CSR/internal/configs"
	"CSR/internal/pkg"
	"CSR/internal/models"
	"errors"

	"strings"

	"github.com/gin-gonic/gin"
)
func (ctrl *Controller)extractTokenFromHeader(c*gin.Context,headerKey string)(string ,error){
	header:=c.GetHeader(headerKey)
	if header==""{
		return "",errors.New("empty request header")
	}
	headerSplit:=strings.Split(header," ")
	if len(headerSplit)!=2{
		return "",errors.New("invalid auth header")	
	}
	if len(headerSplit[1])==0{
		return "",errors.New("empty token error")
	}
	return headerSplit[1],nil
}


func(ctrl*Controller)generateNewTokenPair(userId int,userRole models.Role)(string,string,error){
	accessToken,err:=pkg.GenerateToken(userId,configs.AppSettings.AuthParams.AccessTokenTtl,false,userRole)
	if err!=nil{
		return "","",nil
	}
	refreshToken,err:=pkg.GenerateToken(userId,configs.AppSettings.AuthParams.RefreshTokenTtl,true,userRole)
	if err!=nil{
		return "","",nil
	}
	return accessToken,refreshToken,nil
}