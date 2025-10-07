package contracts

import "CSR/internal/models"

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type ServiceI interface {
	GetAllEmployees() ([]models.Employee, error)
	GetEmployeeById(id int) (models.Employee, error)
	CreateNewEmployee(emp models.EmployeeRequest) error
	UpdateEmployeeById(id int, user models.EmployeeRequest) error
	DeleteEmployeeById(id int) error

	CreateNewUser(userRequest models.SignUpRequest)error
	Authenticate(userRequest models.SignInRequest) (int ,error)
	
}
