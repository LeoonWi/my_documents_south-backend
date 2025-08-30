package repository

import "github.com/jmoiron/sqlx"

type RoleRepository struct {
	Conn *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{Conn: db}
}
