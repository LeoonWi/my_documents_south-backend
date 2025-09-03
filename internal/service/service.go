package service

import (
	"context"
	"errors"
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
	service := model.Service{Name: name}

	if err := s.repositories.Service().CreateService(ctx, &service); err != nil {
		return nil, err
	}

	return &service, nil
}

func (s *ServiceService) GetService(ctx context.Context) *[]model.Service {
	var service []model.Service

	if err := s.repositories.Service().GetService(ctx, &service); err != nil {
		return nil
	}

	return &service
}

func (s *ServiceService) GetServiceById(ctx context.Context, id int) (*model.Service, error) {
	if id < 1 {
		return nil, errors.New("Некорректное значение id")
	}

	service := model.Service{}
	if err := s.repositories.Service().GetServiceById(ctx, id, &service); err != nil {
		return nil, err
	}

	return &service, nil
}
