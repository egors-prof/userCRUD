package Service

import (
	"CSR/Internal/models"
)

func (s *Service) GetAllUsers() ([]models.User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *Service) GetUserById(id int) (models.User, error) {
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *Service) CreateNewUser(user models.User) error {
	err := s.repository.CreateNewUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateUserById(id int, user models.User) error {
	err := s.repository.UpdateUserById(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUserById(id int) error {
	err := s.repository.DeleteUserById(id)
	if err != nil {
		return err
	}
	return nil

}
