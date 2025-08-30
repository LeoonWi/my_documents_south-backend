package service

import (
	"context"
	"my_documents_south_backend/internal/core/repositories"
	"my_documents_south_backend/internal/model"
)

type ServiceService struct {
	repositories repositories.Postgres
}

func NewServiceService(repositories repositories.Postgres) *ServiceService {
	return &ServiceService{repositories: repositories}
}

func (s *ServiceService) CreateService(ctx context.Context, name string) (*model.Service, error) {
	return nil, nil
}
