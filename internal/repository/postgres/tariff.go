package postgres

import (
	"context"
	"my_documents_south_backend/internal/model"

	"github.com/jmoiron/sqlx"
)

type TariffRepository struct {
	Conn *sqlx.DB
}

func NewTariffRepository(db *sqlx.DB) *TariffRepository {
	return &TariffRepository{Conn: db}
}

func (r *TariffRepository) CreateTariff(ctx context.Context, tariff *model.Tariff) error {
	return nil
}

func (r *TariffRepository) GetTariffById(ctx context.Context, id int, tariff *model.Tariff) error {
	return nil
}

func (r *TariffRepository) GetTariffByName(ctx context.Context, name string, tariff *model.Tariff) error {
	return nil
}
