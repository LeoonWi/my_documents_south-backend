package repository

import (
	"context"
	"errors"
	"fmt"
	"my_documents_south_backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type requestRepository struct {
	conn *sqlx.DB
}

func NewRequestRepository(db *sqlx.DB) models.RequestRepository {
	return &requestRepository{conn: db}
}

func (r *requestRepository) Create(c context.Context, req *models.Request) error {
	query := `INSERT INTO "request" (name, service_id, owner_id, employee_id, priority, "desc", status, desired_at)
        	  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        	  RETURNING *`

	// Начинаем транзакцию
	tx, err := r.conn.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	if err := tx.GetContext(
		c,
		req,
		query,
		req.Name,
		req.ServiceId,
		req.OwnerId,
		req.EmployeeId,
		req.Priority,
		req.Desc,
		req.Status,
		req.DesiredAt,
	); err != nil {
		// Отменяем транзакцию, в случае возникнования ошибки
		if rollbackError := tx.Rollback(); rollbackError != nil {
			return fmt.Errorf("failed to rollback transaction: %w", rollbackError)
		}

		// Декрементируем ID до актуальной последней записи
		_, resetErr := r.conn.ExecContext(c, `SELECT setval('request_id_seq', (SELECT COALESCE(MAX(id), 0) FROM "request"))`)
		if resetErr != nil {
			return fmt.Errorf("failed to reset request: %w", resetErr)
		}

		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Применяем транзакцию
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *requestRepository) Get(c context.Context, req *[]models.Request) error { return nil }

func (r *requestRepository) GetById(c context.Context, id int, req *models.Request) error {
	err := r.conn.GetContext(c, req, `SELECT * FROM "request" WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *requestRepository) GetWithFilter(ctx context.Context, req *[]models.Request, filter models.Request) error {
	query := `SELECT * FROM "request" WHERE 1=1`

	args := []interface{}{}
	i := 1

	if filter.OwnerId != 0 {
		query += fmt.Sprintf(" AND owner_id = $%d", i)
		args = append(args, filter.OwnerId)
		i++
	}
	if filter.ServiceId != 0 {
		query += fmt.Sprintf(" AND service_id = $%d", i)
		args = append(args, filter.ServiceId)
		i++
	}
	if !filter.DesiredAt.IsZero() {
		query += fmt.Sprintf(" AND desired_at <= $%d", i)
		args = append(args, filter.DesiredAt)
		i++
	}
	if !filter.DesiredAt.IsZero() {
		query += fmt.Sprintf(" AND status <= $%d", i)
		args = append(args, filter.DesiredAt)
		i++
	}
	if filter.EmployeeId != 0 {
		query += fmt.Sprintf(" AND employee_id = $%d", i)
		args = append(args, filter.EmployeeId)
		i++
	}

	err := r.conn.SelectContext(ctx, req, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *requestRepository) Update(c context.Context, req *models.Request) error { return nil }

func (r *requestRepository) Delete(c context.Context, id int) error {
	result, err := r.conn.ExecContext(c, `DELETE FROM "request" WHERE id=$1`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("request not found")
	}
	return nil
}
