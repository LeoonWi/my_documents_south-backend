package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type ServiceService struct {
	repositories repositories.Postgres
}

func NewServiceService(repositories repositories.Postgres) *ServiceService {
	return &ServiceService{repositories: repositories}
}
