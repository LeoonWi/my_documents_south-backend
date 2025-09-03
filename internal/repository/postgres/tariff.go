package postgres

import (
	"context"
	"errors"
	"my_documents_south_backend/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
)

type TariffRepository struct {
	Conn *sqlx.DB
}

func NewTariffRepository(db *sqlx.DB) *TariffRepository {
	return &TariffRepository{Conn: db}
}

func (r *TariffRepository) CreateTariff(c context.Context, tariff *model.Tariff) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()
	err := r.Conn.GetContext(ctx, tariff, "INSERT INTO tariff (name) VALUES ($1) returning *", tariff.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *TariffRepository) GetTariffs(c context.Context, tariff *[]model.Tariff) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()
	err := r.Conn.SelectContext(ctx, tariff, "SELECT * FROM tariff")
	if err != nil {
		return err
	}
	return nil
}

func (r *TariffRepository) GetTariffByID(c context.Context, id int, tariff *model.Tariff) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()
	err := r.Conn.GetContext(ctx, tariff, "SELECT * FROM tariff WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TariffRepository) GetTariffByName(c context.Context, name string, tariff *model.Tariff) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()
	err := r.Conn.GetContext(ctx, tariff, "SELECT * FROM tariff WHERE name = $1", name) //возможно нужен SelectContext,если допускается повторение имён
	if err != nil {
		return err
	}
	return nil
}

func (r *TariffRepository) UpdateTariff(c context.Context, tariff *model.Tariff) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()
	return r.Conn.GetContext(ctx, tariff, "UPDATE tariff SET name = $1, updated_at = NOW() WHERE id = $2 RETURNING *;", tariff.Name, tariff.Id)
}

func (r *TariffRepository) DeleteTariff(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()
	result, err := r.Conn.ExecContext(ctx, "DELETE FROM tariff WHERE id=$1", id)
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
}
