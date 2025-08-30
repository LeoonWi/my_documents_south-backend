package service

import (
	"my_documents_south_backend/internal/core/repositories"
	"my_documents_south_backend/internal/model"
)

type ServiceService struct {
	repositories repositories.Postgres
}

func NewServiceService(repositories repositories.Postgres) *ServiceService {
	return &ServiceService{repositories: repositories}
}

func (s *ServiceService) CreateService(name string) (*model.Service, error) {
	return nil, nil
}
