package repository

import "github.com/jmoiron/sqlx"

type EmployeeSpecsRepository struct {
	Conn *sqlx.DB
}

func NewEmployeeSpecsRepository(db *sqlx.DB) *EmployeeSpecsRepository {
	return &EmployeeSpecsRepository{Conn: db}
}
