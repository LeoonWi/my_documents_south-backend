package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type EmployeeSpecsService struct {
	repositories repositories.Postgres
}

func NewEmployeeSpecsService(repositories repositories.Postgres) *EmployeeSpecsService {
	return &EmployeeSpecsService{repositories: repositories}
}
