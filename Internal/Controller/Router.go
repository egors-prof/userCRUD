package Controller

func (ctrl *Controller) RegisterEndpoints() {

	ctrl.router.GET("/users", ctrl.GetAllUsers)
	ctrl.router.GET("/users/:id", ctrl.GetUserById)
	ctrl.router.POST("/users", ctrl.CreateNewUser)
	ctrl.router.PUT("/users/:id", ctrl.UpdateUserById)
	ctrl.router.DELETE("/users/:id", ctrl.DeleteUserById)

}
func (ctrl *Controller) RunServer(addr string) error {
	ctrl.RegisterEndpoints()
	if err := ctrl.router.Run(addr); err != nil {
		return err
	}
	return nil
}
