package repository

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

func (r*Repository)GetUserByUsername(userName string)(models.User,error){
	user :=models.User{}
	err:=r.db.Get(&user,`select *
	from users where username =$1`,userName)
	if user.Id==0{
		return models.User{},errs.ErrUserNotFound
	}
	if err!=nil{
		if errors.Is(err,sql.ErrNoRows){
			r.repoLog.Error().Err(sql.ErrNoRows).Send()
			return models.User{} ,errs.ErrUserNotFound
		}

		r.repoLog.Error().Err(err).Msg("unknown error")
		return models.User{},fmt.Errorf("internal error occurred")
	}
	return user,nil
}


func (r *Repository) CreateNewUser(newUser models.SignUpRequest) error {
	r.repoLog.Info().Any("newUser:",newUser).Send()
	user,_:=r.GetUserByUsername(newUser.Username)
	r.repoLog.Info().Any("user",user).Send()
	if user.Id==0{
		r.repoLog.Info().Any("status: ",user).Send()
		_,err:=r.db.Exec(
		`insert into users
		(full_name, username,hash_pass) 
		values 
		($1,$2,$3)`,
		newUser.FullName,
		newUser.Username,
		newUser.Password,
		)
		if err!=nil{
			r.repoLog.Error().Err(err).Msg("error while inserting user")
			return err
		}
		return nil 
	}
	return errs.ErrUserAlreadyExists
}

