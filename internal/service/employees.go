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
func (s *Service) GetAllEmployees() ([]models.Employee, error) {
	emps, err := s.repository.GetAllEmployees()
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return emps, errs.ErrUserNotFound
		}
		return emps, err
	}
	return emps, nil

}
func (s *Service) GetEmployeeById(id int) (models.Employee, error) {
	redis_id := fmt.Sprintf("employee_%d", id)
	emp, err := s.cache.Get(ctx, redis_id)

	if err != nil {
		emp, err := s.repository.GetEmployeeById(id)
		if err != nil {
			return models.Employee{}, errs.ErrUserNotFound
		}
		s.cache.Set(ctx, redis_id, emp,DefaultTTL)
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
	if len(emp.Name) < 4 {
		return errs.ErrInvalidUserName
	}
	err := s.repository.CreateNewEmployee(emp)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateEmployeeById(id int, emp models.EmployeeRequest) error {
	_, err := s.repository.GetEmployeeById(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
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
	_, err := s.repository.GetEmployeeById(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrUserNotFound
		}
	}
	err = s.repository.DeleteEmployeeById(id)
	if err != nil {
		return err
	}
	return nil
}
