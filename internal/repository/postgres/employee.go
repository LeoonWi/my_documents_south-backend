package postgres

import "github.com/jmoiron/sqlx"

type EmployeeRepository struct {
	Conn *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{Conn: db}
}
