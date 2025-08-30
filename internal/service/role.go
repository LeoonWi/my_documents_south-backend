package service

import (
	"my_documents_south_backend/internal/core/repositories"
	"my_documents_south_backend/internal/model"
)

type RoleService struct {
	repositories repositories.Postgres
}

func NewRoleService(repositories repositories.Postgres) *RoleService {
	return &RoleService{repositories: repositories}
}

func (s *RoleService) CreateRole(name string) (*model.Role, error) {
	return nil, nil
}
