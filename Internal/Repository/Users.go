package Repository

import (
	"CSR/Internal/models"
	"errors"
	"log"
)

func (r *Repository) GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	err := r.db.Select(&users, `select id,name,email,age from Users`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return users, nil
}

func (r *Repository) GetUserById(id int) (models.User, error) {
	var checkId int

	err := r.db.Get(&checkId, `select id from users where id =$1`, id)

	if checkId == 0 {
		err := errors.New("no user found with given id")
		log.Println(err)
		return models.User{}, err
	}
	user := models.User{}
	err = r.db.Get(&user, `select id,name,email,age from Users where id=$1`, id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}
func (r *Repository) CreateNewUser(user models.User) error {
	_, err := r.db.Exec(`insert into users (name,email,age)values ($1,$2,$3)`, user.Name, user.Email, user.Age)
	if err != nil {
		log.Println("repo :", err)
		return err
	}
	return nil
}

func (r *Repository) UpdateUserById(id int, user models.User) error {
	var checkId int

	err := r.db.Get(&checkId, `select id from users where id =$1`, id)

	if checkId == 0 {
		err := errors.New("no user found with given id")
		log.Println(err)
		return err
	}
	_, err = r.db.Exec(`update users set name =$1,email=$2,age=$3 where id=$4`, user.Name, user.Email, user.Age, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteUserById(id int) error {
	var checkId int

	err := r.db.Get(&checkId, `select id from users where id =$1`, id)

	if checkId == 0 {
		err := errors.New("no user found with given id")
		log.Println(err)
		return err
	}

	_, err = r.db.Exec(`delete from users where id =$1`, id)
	if err != nil {
		log.Println("error while on repository side while deleting")
		return err
	}
	return nil
}
