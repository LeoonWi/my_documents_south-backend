package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type RoleService struct {
	repositories repositories.Postgres
}

func NewRoleService(repositories repositories.Postgres) *RoleService {
	return &RoleService{repositories: repositories}
}
