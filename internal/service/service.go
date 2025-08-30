package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type ServiceService struct {
	repositories repositories.Repositories
}

func NewServiceService(repositories repositories.Repositories) *ServiceService {
	return &ServiceService{repositories: repositories}
}
