package controller

import (
	"CSR/internal/configs"
	_ "CSR/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (ctrl *Controller) RegisterEndpoints() {

	ctrl.router.GET("/users", ctrl.GetAllUsers)
	ctrl.router.GET("/users/:id", ctrl.GetUserById)
	ctrl.router.POST("/users", ctrl.CreateNewUser)
	ctrl.router.PUT("/users/:id", ctrl.UpdateUserById)
	ctrl.router.DELETE("/users/:id", ctrl.DeleteUserById)
	ctrl.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
func (ctrl *Controller) RunServer() error {
	ctrl.RegisterEndpoints()
	if err := ctrl.router.Run(configs.AppSettings.AppParam.Port); err != nil {
		return err
	}
	return nil
}
