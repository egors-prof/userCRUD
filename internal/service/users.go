package service

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
	
)
var ctx = context.Background()
var DefaultTTL=3*time.Minute
func (s *Service) GetAllUsers() ([]models.User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return users, errs.ErrUserNotFound
		}
		return users, err
	}
	return users, nil

}
func (s *Service) GetUserById(id int) (models.User, error) {
	redis_id := fmt.Sprintf("user_%d", id)
	user, err := s.cache.Get(ctx, redis_id)

	if err != nil {
		user, err := s.repository.GetUserById(id)
		log.Println("service : ", user, err)

		if err != nil {
			return models.User{}, errs.ErrUserNotFound
		}
		s.cache.Set(ctx, redis_id, user,DefaultTTL)
		log.Printf("user_%d is cached", id)
		return user, nil
	}
	userRes := models.User{}
	err = json.Unmarshal([]byte(user), &userRes)
	if err != nil {
		return models.User{}, err
	}
	return userRes, nil
}

func (s *Service) CreateNewUser(user models.User) error {
	if len(user.Name) < 4 {
		return errs.ErrInvalidUserName
	}
	err := s.repository.CreateNewUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateUserById(id int, user models.User) error {
	_, err := s.repository.GetUserById(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	err = s.repository.UpdateUserById(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUserById(id int) error {
	_, err := s.repository.GetUserById(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrUserNotFound
		}
	}
	err = s.repository.DeleteUserById(id)
	if err != nil {
		return err
	}
	return nil
}
