package postgres

import (
	"context"
	"my_documents_south_backend/internal/model"

	"github.com/jmoiron/sqlx"
)

type RoleRepository struct {
	Conn *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{Conn: db}
}

func (r *RoleRepository) CreateRole(ctx context.Context, role *model.Role) error {
	return nil
}

func (r *RoleRepository) GetRoleByName(ctx context.Context, name string, role *model.Role) error {
	return nil
}
