package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"my_documents_south_backend/internal/models"
	"time"
)

type serviceRepository struct {
	conn *sqlx.DB
}

func NewServiceRepository(db *sqlx.DB) models.ServiceRepository {
	return &serviceRepository{conn: db}
}

func (r *serviceRepository) Create(c context.Context, service *models.Service) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	err := r.conn.GetContext(ctx, service, "INSERT INTO service (name) VALUES ($1) RETURNING *", service.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *serviceRepository) Get(c context.Context, service *[]models.Service) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	if err := r.conn.SelectContext(ctx, service, "SELECT * FROM service"); err != nil {
		return err
	}
	return nil
}

func (r *serviceRepository) GetById(c context.Context, id int, service *models.Service) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	err := r.conn.GetContext(ctx, service, "SELECT * FROM service WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *serviceRepository) Update(c context.Context, service *models.Service) error {
	// TODO update service repository
	return nil
}

func (r *serviceRepository) Delete(c context.Context, id int) error {
	// TODO delete service repository
	return nil
}
