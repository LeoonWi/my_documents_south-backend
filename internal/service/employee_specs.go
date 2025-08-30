package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type EmployeeSpecsService struct {
	repositories repositories.Repositories
}

func NewEmployeeSpecsService(repositories repositories.Repositories) *EmployeeSpecsService {
	return &EmployeeSpecsService{repositories: repositories}
}
