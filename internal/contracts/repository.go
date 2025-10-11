package contracts

import (
	"CSR/internal/models"
)

type RepositoryI interface {
	GetAllEmployees() ([]models.Employee, error)
	GetEmployeeById(id int) (models.Employee, error)
	CreateNewEmployee(employee models.EmployeeRequest) error
	UpdateEmployeeById(id int, user models.EmployeeRequest) error
	DeleteEmployeeById(id int) error
	CreateNewUser(userRequest models.SignUpRequest) error
	GetUserByUsername(userName string) (models.User, error)
}
