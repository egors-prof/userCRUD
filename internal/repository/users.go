package repository

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"context"

	// "encoding/json"

	// "errors"
	// "fmt"
	
	"time"
	// "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func (r *Repository) GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	time.Sleep(10 * time.Second)
	err := r.db.Select(&users, `select id,name,email,age from Users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) GetUserById(id int) (models.User, error) {
	user:=models.User{}
	err:=r.db.Get(&user, `select id,name,email,age from users where id=$1`,id)
	if user.Id==0{
		return models.User{},errs.ErrUserNotFound
	}
	if err!=nil{
		return models.User{},err
	}
	return user, err
}

func (r *Repository) CreateNewUser(user models.User) error {
	_, err := r.db.Exec(`insert into users (name,email,age)values ($1,$2,$3)`, user.Name, user.Email, user.Age)
	if err != nil {
		return r.transferError(err)
	}
	return nil
}

func (r *Repository) UpdateUserById(id int, user models.User) error {
	_, err := r.db.Exec(`update users set name =$1,email=$2,age=$3 where id=$4`, user.Name, user.Email, user.Age, id)
	if err != nil {
		return r.transferError(err)
	}
	return nil
}

func (r *Repository) DeleteUserById(id int) error {
	_, err := r.db.Exec(`delete from users where id =$1`, id)
	if err != nil {
		return r.transferError(err)
	}
	return nil
}
