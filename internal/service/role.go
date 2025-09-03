package service

import (
	"context"
	"errors"
	"my_documents_south_backend/internal/core/repositories"
	"my_documents_south_backend/internal/model"
)

type RoleService struct {
	repositories repositories.Postgres
}

func NewRoleService(repositories repositories.Postgres) *RoleService {
	return &RoleService{repositories: repositories}
}

func (s *RoleService) CreateRole(ctx context.Context, name string) (*model.Role, error) {
	role := model.Role{Name: name}

	if err := s.repositories.Role().CreateRole(ctx, &role); err != nil {
		return nil, err
	}

	return &role, nil
}

func (s *RoleService) GetRoles(ctx context.Context) *[]model.Role {
	var roles []model.Role

	if err := s.repositories.Role().GetRoles(ctx, &roles); err != nil {
		return nil
	}

	return &roles
}

func (s *RoleService) GetRoleById(ctx context.Context, id int) (*model.Role, error) {
	if id < 1 {
		return nil, errors.New("Некорректное значение id")
	}

	role := model.Role{}
	if err := s.repositories.Role().GetRoleById(ctx, id, &role); err != nil {
		return nil, err
	}

	return &role, nil
}
