package repository

import "github.com/jmoiron/sqlx"

type TariffRepository struct {
	Conn *sqlx.DB
}

func NewTariffRepository(db *sqlx.DB) *TariffRepository {
	return &TariffRepository{Conn: db}
}
