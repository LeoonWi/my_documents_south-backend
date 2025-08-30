package role

import "github.com/jmoiron/sqlx"

type Repository struct {
	Conn *sqlx.DB
}
