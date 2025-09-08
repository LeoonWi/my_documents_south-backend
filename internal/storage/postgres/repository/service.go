package repository

import (
	"context"
	"errors"
	"my_documents_south_backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type serviceRepository struct {
	conn *sqlx.DB
}

func NewServiceRepository(db *sqlx.DB) models.ServiceRepository {
	return &serviceRepository{conn: db}
}

func (r *serviceRepository) Create(c context.Context, service *models.Service) error {
	err := r.conn.GetContext(c, service, "INSERT INTO service (name) VALUES ($1) RETURNING *", service.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *serviceRepository) Get(c context.Context, service *[]models.Service) error {
	if err := r.conn.SelectContext(c, service, "SELECT * FROM service"); err != nil {
		return err
	}
	return nil
}

func (r *serviceRepository) GetById(c context.Context, id int, service *models.Service) error {
	err := r.conn.GetContext(c, service, "SELECT * FROM service WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *serviceRepository) Update(c context.Context, service *models.Service) error {
	return r.conn.GetContext(c, service, "UPDATE service SET name = $1, update_at = NOW() WHERE id = $2 RETURNING *;", service.Name, service.Id)
}

func (r *serviceRepository) Delete(c context.Context, id int) error {
	result, err := r.conn.ExecContext(c, "DELETE FROM service WHERE id=$1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("service not found")
	}

	return nil
}
