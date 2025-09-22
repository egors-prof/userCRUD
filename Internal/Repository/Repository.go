package Repository

import (
	"CSR/Internal/errs"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) transferError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Println(err)
		return errs.ErrNotFound
	default:
		return err
	}
}
