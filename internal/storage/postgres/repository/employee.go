package repository

import (
	"context"
	"my_documents_south_backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type employeeRepository struct {
	conn *sqlx.DB
}

func NewEmployeeRepository(conn *sqlx.DB) models.EmployeeRepository {
	return &employeeRepository{conn: conn}
}

func (r *employeeRepository) Create(c context.Context, employee *models.Employee) error {
	return nil
}

func (r *employeeRepository) Get(c context.Context, employee *[]models.Employee) error {
	return nil
}

func (r *employeeRepository) GetById(c context.Context, id int, employee *models.Employee) error {
	return nil
}

func (r *employeeRepository) Update(c context.Context, employee *models.Employee) error {
	return nil
}

func (r *employeeRepository) Delete(c context.Context, id int) error {
	return nil
}
