package postgres

import (
	"context"
	"my_documents_south_backend/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
)

type ServiceRepository struct {
	Conn *sqlx.DB
}

func NewServiceRepository(db *sqlx.DB) *ServiceRepository {
	return &ServiceRepository{Conn: db}
}

func (r *ServiceRepository) CreateService(c context.Context, service *model.Service) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	err := r.Conn.GetContext(ctx, service, "INSERT INTO service (name) VALUES ($1) RETURNING *", service.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *ServiceRepository) GetService(c context.Context, service *[]model.Service) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	if err := r.Conn.SelectContext(ctx, service, "SELECT * FROM service"); err != nil {
		return err
	}
	return nil
}

func (r *ServiceRepository) GetServiceById(c context.Context, id int, service *model.Service) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	err := r.Conn.GetContext(ctx, service, "SELECT * FROM service WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ServiceRepository) GetServiceByName(ctx context.Context, name string, service *model.Service) error {
	return nil
}
