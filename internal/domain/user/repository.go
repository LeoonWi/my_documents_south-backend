package user

import "github.com/jmoiron/sqlx"

type Repository struct {
	DB *sqlx.DB
}
