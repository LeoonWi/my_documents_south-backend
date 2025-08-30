package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type EmployeeService struct {
	repositories repositories.Repositories
}

func NewEmployeeService(repositories repositories.Repositories) *EmployeeService {
	return &EmployeeService{repositories: repositories}
}
