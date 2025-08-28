package postgres

import "github.com/jmoiron/sqlx"

type TariffRepository struct {
	db *sqlx.DB
}
