package controller

import (
	_ "CSR/docs"
	"CSR/internal/configs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (ctrl *Controller) RegisterEndpoints() {

	{
apiGroup:=ctrl.router.Group("/api",ctrl.checkUserAuthentication)
	apiGroup.GET("/employees", ctrl.GetAllEmployees)
	apiGroup.GET("/employees/:id", ctrl.GetEmployeeById)
	apiGroup.POST("/employees", ctrl.CreateNewEmployee)
	apiGroup.PUT("/employees/:id", ctrl.UpdateUserById)
	apiGroup.DELETE("/employees/:id", ctrl.DeleteUserById)
	}
	
	ctrl.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
{
	authGroup:=ctrl.router.Group("/auth")
	authGroup.POST("/sign-up",ctrl.SignUp)
	authGroup.POST("/sign-in",ctrl.SignIn)
	authGroup.GET("/refresh",ctrl.RefreshTokenPair)
	authGroup.POST("/check",ctrl.checkToken)

}
	
}
func (ctrl *Controller) RunServer() error {
	ctrl.RegisterEndpoints()
	if err := ctrl.router.Run(configs.AppSettings.AppParam.Port); err != nil {
		return err
	}
	return nil
}
