package postgres

import "github.com/jmoiron/sqlx"

type RoleRepository struct {
	db *sqlx.DB
}
