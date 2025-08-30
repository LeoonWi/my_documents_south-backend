package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type RoleService struct {
	repositories repositories.Repositories
}

func NewRoleService(repositories repositories.Repositories) *RoleService {
	return &RoleService{repositories: repositories}
}
