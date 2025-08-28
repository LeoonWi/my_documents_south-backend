package postgres

import "github.com/jmoiron/sqlx"

type SvcRepository struct {
	db *sqlx.DB
}
