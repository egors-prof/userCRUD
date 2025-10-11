package service

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"log"
	"os"
	"time"
)

var ctx = context.Background()
var DefaultTTL = 3 * time.Minute

func (s *Service) GetAllEmployees() ([]models.Employee, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "service.GetAllEmployees").Logger()
	emps, err := s.repository.GetAllEmployees()
	if err != nil {
		logger.Error().Err(err).Send()
		if errors.Is(err, errs.ErrNotFound) {
			return emps, errs.ErrUserNotFound
		}
		return emps, err
	}
	return emps, nil

}
func (s *Service) GetEmployeeById(id int) (models.Employee, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "service.GetEmployeeById").Logger()
	redis_id := fmt.Sprintf("employee_%d", id)
	emp, err := s.cache.Get(ctx, redis_id)

	if err != nil {
		emp, err := s.repository.GetEmployeeById(id)
		if err != nil {
			logger.Error().Err(errs.ErrUserNotFound).Send()
			return models.Employee{}, errs.ErrUserNotFound
		}
		s.cache.Set(ctx, redis_id, emp, DefaultTTL)
		log.Printf("employee_%d is cached", id)
		return emp, nil
	}
	empRes := models.Employee{}
	err = json.Unmarshal([]byte(emp), &empRes)
	if err != nil {
		return models.Employee{}, err
	}
	return empRes, nil
}

func (s *Service) CreateNewEmployee(emp models.EmployeeRequest) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "service.CreateNewNewEmployee").Logger()
	if len(emp.Name) < 4 {
		return errs.ErrInvalidUserName
	}
	err := s.repository.CreateNewEmployee(emp)
	if err != nil {
		logger.Error().Err(err).Send()
		return err
	}
	return nil
}

func (s *Service) UpdateEmployeeById(id int, emp models.EmployeeRequest) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "service.UpdateEmployeeById").Logger()
	_, err := s.repository.GetEmployeeById(id)
	if err != nil {
		logger.Error().Err(err).Send()
		if errors.Is(err, errs.ErrNotFound) {
			logger.Error().Err(err).Send()
			return errs.ErrUserNotFound
		}
		return err
	}
	err = s.repository.UpdateEmployeeById(id, emp)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteEmployeeById(id int) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "service.DeleteEmployeeById").Logger()
	_, err := s.repository.GetEmployeeById(id)
	if err != nil {
		logger.Error().Err(err).Send()
		if errors.Is(err, errs.ErrNotFound) {
			logger.Error().Err(err).Send()
			return errs.ErrUserNotFound
		}
	}
	err = s.repository.DeleteEmployeeById(id)
	if err != nil {
		logger.Error().Err(err).Send()
		return err
	}
	return nil
}
