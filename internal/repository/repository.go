package repository

import (
	"CSR/internal/errs"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Repository struct {
	db    *sqlx.DB
	Cache *Cache
	repoLog zerolog.Logger
}

func NewRepository(db *sqlx.DB, cache *Cache,logger zerolog.Logger) *Repository {
	return &Repository{
		db: db,
		Cache:cache,
		repoLog:logger,
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
