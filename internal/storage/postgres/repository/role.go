package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"my_documents_south_backend/internal/models"
	"time"
)

type roleRepository struct {
	conn *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) models.RoleRepository {
	return &roleRepository{conn: db}
}

func (r *roleRepository) Create(c context.Context, role *models.Role) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	err := r.conn.GetContext(ctx, role, "INSERT INTO role (name) VALUES ($1) RETURNING *", role.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) Get(c context.Context, roles *[]models.Role) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	if err := r.conn.SelectContext(ctx, roles, "SELECT * FROM role"); err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) GetById(c context.Context, id int, role *models.Role) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	err := r.conn.GetContext(ctx, role, "SELECT * FROM role WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) Update(c context.Context, role *models.Role) error {
	// TODO update role repository
	return nil
}

func (r *roleRepository) Delete(c context.Context, id int) error {
	// TODO delete role repository
	return nil
}
