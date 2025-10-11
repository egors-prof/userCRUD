package repository

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"context"
	"github.com/rs/zerolog"
	"os"
)

var ctx = context.Background()

func (r *Repository) GetAllEmployees() ([]models.Employee, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetAllEmployees").Logger()
	emps := []models.Employee{}
	err := r.db.Select(&emps, `select id,name,email,age from Employees`)
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, r.transferError(err)
	}
	return emps, nil
}

func (r *Repository) GetEmployeeById(id int) (models.Employee, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetEmployeeById").Logger()
	emp := models.Employee{}
	err := r.db.Get(&emp, `select id,name,email,age from Employees where id=$1`, id)
	if emp.ID == 0 {
		logger.Error().Err(err).Send()
		return models.Employee{}, errs.ErrUserNotFound
	}
	if err != nil {
		logger.Error().Err(err).Send()
		r.transferError(err)
	}

	if err != nil {
		logger.Error().Err(err).Send()
		return models.Employee{}, r.transferError(err)
	}
	return emp, nil
}

func (r *Repository) CreateNewEmployee(employee models.EmployeeRequest) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateNewEmployee").Logger()
	_, err := r.db.Exec(`insert into Employees (name,email,age)values ($1,$2,$3)`, employee.Name, employee.Email, employee.Age)
	if err != nil {
		logger.Error().Err(err).Send()
		return r.transferError(err)
	}
	return nil
}

func (r *Repository) UpdateEmployeeById(id int, employee models.EmployeeRequest) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateNewEmployee").Logger()
	_, err := r.db.Exec(`update Employees set name =$1,email=$2,age=$3 where id=$4`, employee.Name, employee.Email, employee.Age, id)
	if err != nil {
		logger.Error().Err(err).Send()
		return r.transferError(err)
	}
	return nil
}

func (r *Repository) DeleteEmployeeById(id int) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateNewEmployee").Logger()
	_, err := r.db.Exec(`delete from Employees where id =$1`, id)
	if err != nil {
		logger.Error().Err(err).Send()
		return r.transferError(err)
	}
	return nil
}
