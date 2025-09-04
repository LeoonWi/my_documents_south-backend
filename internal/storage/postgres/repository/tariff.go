package repository

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"my_documents_south_backend/internal/models"
)

type tariffRepository struct {
	conn *sqlx.DB
}

func NewTariffRepository(db *sqlx.DB) models.TariffRepository {
	return &tariffRepository{conn: db}
}

func (r *tariffRepository) Create(c context.Context, tariff *models.Tariff) error {
	err := r.conn.GetContext(c, tariff, "INSERT INTO tariff (name) VALUES ($1) returning *", tariff.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *tariffRepository) Get(c context.Context, tariff *[]models.Tariff) error {
	err := r.conn.SelectContext(c, tariff, "SELECT * FROM tariff")
	if err != nil {
		return err
	}
	return nil
}

func (r *tariffRepository) GetById(c context.Context, id int, tariff *models.Tariff) error {
	err := r.conn.GetContext(c, tariff, "SELECT * FROM tariff WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *tariffRepository) Update(c context.Context, tariff *models.Tariff) error {
	return r.conn.GetContext(c, tariff, "UPDATE tariff SET name = $1, updated_at = NOW() WHERE id = $2 RETURNING *;", tariff.Name, tariff.Id)
}

func (r *tariffRepository) Delete(c context.Context, id int) error {
	result, err := r.conn.ExecContext(c, "DELETE FROM tariff WHERE id=$1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("tariff not found")
	}
	return nil
}
