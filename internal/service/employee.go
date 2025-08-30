package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type EmployeeService struct {
	repositories repositories.Postgres
}

func NewEmployeeService(repositories repositories.Postgres) *EmployeeService {
	return &EmployeeService{repositories: repositories}
}
