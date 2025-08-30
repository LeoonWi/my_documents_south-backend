package tariff

import "github.com/jmoiron/sqlx"

type Repository struct {
	Conn *sqlx.DB
}
