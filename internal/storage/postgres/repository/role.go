package repository

import (
	"context"
	"errors"
	"my_documents_south_backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type roleRepository struct {
	conn *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) models.RoleRepository {
	return &roleRepository{conn: db}
}

func (r *roleRepository) Create(c context.Context, role *models.Role) error {
	err := r.conn.GetContext(c, role, "INSERT INTO role (name) VALUES ($1) RETURNING *", role.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) Get(c context.Context, roles *[]models.Role) error {
	if err := r.conn.SelectContext(c, roles, "SELECT * FROM role"); err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) GetById(c context.Context, id int, role *models.Role) error {
	err := r.conn.GetContext(c, role, "SELECT * FROM role WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) SetSuperRole(c context.Context, id int) error {
	var count int
	err := r.conn.GetContext(c, &count, `SELECT COUNT(*) FROM setting`)
	if err != nil {
		return err
	}

	if count != 0 {
		_, err = r.conn.ExecContext(c, `UPDATE setting SET superuser_role_id = $1`, id)
		if err != nil {
			return err
		}
	}

	_, err = r.conn.ExecContext(c, `INSERT INTO setting (superuser_role_id) VALUES ($1)`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) GetSuperRole(c context.Context, role *models.Role) error {
	err := r.conn.GetContext(c, role, `SELECT * FROM "role" r WHERE r.id = setting.superuser_role_id`)
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) Update(c context.Context, role *models.Role) error {
	return r.conn.GetContext(c, role, "UPDATE role SET name = $1, updated_at = NOW() WHERE id = $2 RETURNING *;", role.Name, role.Id)
}

func (r *roleRepository) Delete(c context.Context, id int) error {
	result, err := r.conn.ExecContext(c, "DELETE FROM role WHERE id=$1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("role not found")
	}
	return nil
}
