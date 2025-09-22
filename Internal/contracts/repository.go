package contracts

import "CSR/Internal/models"

type RepositoryI interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	CreateNewUser(user models.User) error
	UpdateUserById(id int, user models.User) error
	DeleteUserById(id int) error
}
