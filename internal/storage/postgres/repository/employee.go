package repository

import (
	"context"
	"fmt"
	"my_documents_south_backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type employeeRepository struct {
	conn *sqlx.DB
}

// NewEmployeeRepository return pointer by employeeRepository
func NewEmployeeRepository(conn *sqlx.DB) models.EmployeeRepository {
	return &employeeRepository{conn: conn}
}

func (r *employeeRepository) Create(c context.Context, employee *models.Employee) error {
	query := `INSERT INTO "employee" (name, last_name, middle_name, email, password, role_id, active)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)
			  RETURNING *`

	// Начинаем транзакцию
	tx, err := r.conn.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Выполняем запрос на добавление сотрудника
	if err := tx.GetContext(
		c,
		employee,
		query,
		employee.Name,
		employee.LastName,
		employee.MiddleName,
		employee.Email,
		employee.Password,
		employee.RoleId,
		true,
	); err != nil {
		// Отменяем транзакцию, в случае возникнования ошибки
		if rollbackError := tx.Rollback(); rollbackError != nil {
			return fmt.Errorf("failed to rollback transaction: %w", rollbackError)
		}
		// Декрементируем ID до актуальной последней записи
		_, resetErr := r.conn.ExecContext(c, `SELECT setval('employee_id_seq', (SELECT COALESCE(MAX(id), 0) FROM employee))`)
		if resetErr != nil {
			return fmt.Errorf("failed to reset employee: %w", resetErr)
		}

		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Применяем транзакцию
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *employeeRepository) Get(c context.Context, employee *[]models.Employee) error {
	query := `SELECT
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
			LEFT JOIN "role" r ON e.role_id = r.id`

	err := r.conn.SelectContext(c, employee, query)
	if err != nil {
		return fmt.Errorf("failed to get employee: %w", err)
	}
	return nil
}

func (r *employeeRepository) GetById(c context.Context, id int, employee *models.Employee) error {
	err := r.conn.GetContext(c, employee, `SELECT * FROM "user" WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to get employee: %w", err)
	}
	return nil
}

func (r *employeeRepository) GetByEmail(c context.Context, email string, employee *models.Employee) error {
	err := r.conn.GetContext(c, employee, `SELECT * FROM "employee" WHERE "email" = $1`, email)
	if err != nil {
		return fmt.Errorf("not fount employee by %s", email)
	}

	return nil
}

// Update TODO
func (r *employeeRepository) Update(c context.Context, employee *models.Employee) error { return nil }

func (r *employeeRepository) Delete(c context.Context, id int) error {
	_, err := r.conn.ExecContext(c, `DELETE FROM "employee" WHERE id=$1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %w", err)
	}
	return nil
}
