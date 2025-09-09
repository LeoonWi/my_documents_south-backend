package repository

import (
	"context"
	"errors"
	"fmt"
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
	query := `
		INSERT INTO "employee" (id, name, last_name, middle_name, email, password, role_id, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING *
	`

	err := r.conn.GetContext(
		c,
		employee,
		query,
		employee.Name,
		employee.LastName,
		employee.MiddleName,
		employee.Email,
		employee.RoleId,
		employee.Active,
		employee.CreatedAt,
		employee.UpdatedAt,
	)

	if err != nil {
		_, seqErr := r.conn.ExecContext(c, `SELECT setval('user_id_seq', COALESCE(MAX(id), 1)) FROM "user";`)
		if seqErr != nil {
			return fmt.Errorf("failed to reset sequence after error: %w", seqErr)
		}
		return err
	}
	return nil
}

func (r *employeeRepository) Get(c context.Context, employee *[]models.Employee) error {
	err := r.conn.SelectContext(
		c,
		employee,
		`SELECT
			e.id,
			e.name,
			e.last_name,
			e.middle_name,
			e.email,
			e.password,
			e.active,
			e.created_at,
			e.updated_at,
			r.id AS "role.id",
			r.name AS "role.name"
		FROM "employee" e
		LEFT JOIN "role" r ON e.role_id = r.id`,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *employeeRepository) GetById(c context.Context, id int, employee *models.Employee) error {
	err := r.conn.GetContext(c, employee, `SELECT * FROM "user" WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *employeeRepository) Update(c context.Context, employee *models.Employee) error {
	return nil
}

func (r *employeeRepository) Delete(c context.Context, id int) error {
	result, err := r.conn.ExecContext(c, `DELETE FROM "employee" WHERE id=$1`, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
