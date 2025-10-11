package repository

import (
	"CSR/internal/errs"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db    *sqlx.DB
	Cache *Cache
}

func NewRepository(db *sqlx.DB, cache *Cache) *Repository {
	return &Repository{
		db:    db,
		Cache: cache,
	}
}

func (r *Repository) transferError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotFound
	default:
		return err
	}
}
