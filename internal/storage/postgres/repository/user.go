package repository

import (
	"context"
	"errors"
	"fmt"
	"my_documents_south_backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) models.UserRepository {
	return &userRepository{conn: conn}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO "user" (name, last_name, middle_name, email, phone, password, tariff_id, inn, snils)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING *
    `

	err := r.conn.GetContext(
		ctx,
		user,
		query,
		user.Name,
		user.LastName,
		user.MiddleName,
		user.Email,
		user.Phone,
		user.Password,
		user.TariffId,
		user.Inn,
		user.Snils,
	)

	if err != nil {
		_, seqErr := r.conn.ExecContext(ctx, `SELECT setval('user_id_seq', COALESCE(MAX(id), 1)) FROM "user";`)
		if seqErr != nil {
			return fmt.Errorf("failed to reset sequence after error: %w", seqErr)
		}
		return err
	}
	return nil
}

func (r *userRepository) Get(c context.Context, user *[]models.User) error {
	// TODO не работает запрос. Переход на SQLBoiler
	err := r.conn.SelectContext(
		c,
		user,
		`SELECT 
			u.id,
			u.name,
			u.last_name,
			u.middle_name,
			u.email,
			u.phone,
			u.inn,
			u.snils,
			u.created_at,
			u.updated_at,
			t.id AS "tariff.id",
			t.name AS "tariff.name"
		FROM "user" u
		LEFT JOIN "tariff" t ON u.tariff_id = t.id`,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetById(c context.Context, id int, user *models.User) error {
	err := r.conn.GetContext(c, user, `SELECT * FROM "user" WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Update(c context.Context, user *models.User) error {
	// TODO update user repository
	// DONT TOUCH
	return nil
}

func (r *userRepository) Delete(c context.Context, id int) error {
	result, err := r.conn.ExecContext(c, `DELETE FROM "user" WHERE id=$1`, id)
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
func (r *userRepository) UpdateTariffForUsers(ctx context.Context, oldTariffId, newTariffId int) error {
	_, err := r.conn.ExecContext(ctx, `UPDATE "user" SET tariff_id = $1 WHERE tariff_id = $2`, newTariffId, oldTariffId)
	return err
}
