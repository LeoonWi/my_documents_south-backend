package postgres

import (
	"context"
	"my_documents_south_backend/internal/model"

	"github.com/jmoiron/sqlx"
)

type ServiceRepository struct {
	Conn *sqlx.DB
}

func NewServiceRepository(db *sqlx.DB) *ServiceRepository {
	return &ServiceRepository{Conn: db}
}

func (r *ServiceRepository) CreateService(ctx context.Context, service *model.Service) error {
	return nil
}
func (r *ServiceRepository) GetServiceByName(ctx context.Context, name string, service *model.Service) error {
	return nil
}
