package postgres

import "github.com/jmoiron/sqlx"

type ServiceRepository struct {
	Conn *sqlx.DB
}

func NewServiceRepository(db *sqlx.DB) *ServiceRepository {
	return &ServiceRepository{Conn: db}
}
