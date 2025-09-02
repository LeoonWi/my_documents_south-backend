package postgres

import (
	"context"
	"my_documents_south_backend/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
)

type RoleRepository struct {
	Conn *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{Conn: db}
}

func (r *RoleRepository) CreateRole(c context.Context, role *model.Role) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	err := r.Conn.GetContext(ctx, role, "INSERT INTO role (name) VALUES ($1) RETURNING *", role.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *RoleRepository) GetRoles(c context.Context, roles *[]model.Role) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	if err := r.Conn.SelectContext(ctx, roles, "SELECT * FROM role"); err != nil {
		return err
	}

	return nil
}

func (r *RoleRepository) GetRoleById(c context.Context, id int, role *model.Role) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	err := r.Conn.GetContext(ctx, role, "SELECT * FROM role WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *RoleRepository) GetRoleByName(ctx context.Context, name string, role *model.Role) error {
	return nil
}
