package Controller

import (
	"CSR/Internal/Service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service.Service
	router  *gin.Engine
}

func NewController(service *Service.Service) *Controller {
	return &Controller{service: service, router: gin.Default()}
}
