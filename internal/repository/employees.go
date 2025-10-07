package repository

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"context"

	
)

var ctx = context.Background()

func (r *Repository) GetAllEmployees() ([]models.Employee, error) {
	r.repoLog.Info().Str("func_name","GetAllEmployees").Send()
	emps := []models.Employee{}
	err := r.db.Select(&emps, `select id,name,email,age from Employees`)
	if err != nil {
		r.repoLog.Error().Err(err).Msg("error while select")
		return nil, r.transferError(err)
	}
	return emps, nil
}

func (r *Repository) GetEmployeeById(id int) (models.Employee, error) {
	r.repoLog.Info().Str("func_name","GetEmployeeById").Send()
	emp := models.Employee{}
	err := r.db.Get(&emp, `select id,name,email,age from Employees where id=$1`, id)
	if emp.ID == 0 {
		r.repoLog.Error().Err(err).Msg("zero id error")

		return models.Employee{}, errs.ErrUserNotFound
	}
	if err!=nil{
		r.transferError(err)
	}
	
	if err != nil {
		r.repoLog.Error().Err(err).Msg("get request error")

		return models.Employee{}, r.transferError(err)
	}
	return emp, nil
}

func (r *Repository) CreateNewEmployee(employee models.EmployeeRequest) error {
	r.repoLog.Info().Str("func_name"," CreateNewEmployee").Send()
	_, err := r.db.Exec(`insert into Employees (name,email,age)values ($1,$2,$3)`, employee.Name, employee.Email, employee.Age)
	if err != nil {
		r.repoLog.Error().Err(err).Msg("error while creating user")
		return r.transferError(err)
	}
	return nil
}

func (r *Repository) UpdateEmployeeById(id int, employee models.EmployeeRequest) error {
	r.repoLog.Info().Str("func_name","UpdateEmployeeById").Send()
	_, err := r.db.Exec(`update Employees set name =$1,email=$2,age=$3 where id=$4`, employee.Name, employee.Email, employee.Age, id)
	if err != nil {
		r.repoLog.Error().Err(err).Msg("error while updating user")
		return r.transferError(err)
	}
	return nil
}

func (r *Repository) DeleteEmployeeById(id int) error {
	r.repoLog.Info().Str("func_name","DeleteEmployeeById").Send()
	_, err := r.db.Exec(`delete from Employees where id =$1`, id)
	if err != nil {
		r.repoLog.Error().Err(err).Msg("error while deleting user")
		return r.transferError(err)
	}
	return nil
}
