package repository

import (
	"CSR/internal/errs"
	"CSR/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

func (r *Repository) GetUserByUsername(userName string) (models.User, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetUserByUsername").Logger()
	user := models.User{}
	err := r.db.Get(&user, `select *
	from users where username =$1`, userName)
	logger.Debug().Any("user",user).Send()
	if user.Id == 0 {
		return models.User{}, errs.ErrUserNotFound
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error().Err(err).Send()
			return models.User{}, errs.ErrUserNotFound
		}

		logger.Error().Err(err).Send()
		return models.User{}, fmt.Errorf("internal error occurred")
	}
	return user, nil
}

func (r *Repository) CreateNewUser(newUser models.SignUpRequest) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetUserByUsername").Logger()
	user, _ := r.GetUserByUsername(newUser.Username)
	logger.Info().Any("newUser",newUser).Send()
	logger.Info().Any("user",user).Send()
	if user.Id == 0 {
		_, err := r.db.Exec(
			`insert into users
		(full_name, username,hash_pass) 
		values 
		($1,$2,$3)`,
			newUser.FullName,
			newUser.Username,
			newUser.Password,
		)
		if err != nil {
			logger.Error().Err(err).Send()
			return err
		}
		return nil
	}
	return errs.ErrUserAlreadyExists
}
