package repository

import "github.com/jmoiron/sqlx"

type RequestRepository struct {
	Conn *sqlx.DB
}

func NewRequestRepository(db *sqlx.DB) *RequestRepository {
	return &RequestRepository{Conn: db}
}
