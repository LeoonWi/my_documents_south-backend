package services

import (
	"context"
	"errors"
	"my_documents_south_backend/internal/models"
	"time"
)

type serviceService struct {
	serviceRepository models.ServiceRepository
	contextTimeout    time.Duration
}

func NewServiceService(serviceRepository models.ServiceRepository, contextTimeout time.Duration) models.ServiceService {
	return &serviceService{serviceRepository: serviceRepository, contextTimeout: contextTimeout}
}

func (s *serviceService) Create(ctx context.Context, name string) (*models.Service, error) {
	service := models.Service{Name: name}

	if err := s.serviceRepository.Create(ctx, &service); err != nil {
		return nil, err
	}

	return &service, nil
}

func (s *serviceService) Get(ctx context.Context) *[]models.Service {
	var service []models.Service

	if err := s.serviceRepository.Get(ctx, &service); err != nil {
		return nil
	}

	return &service
}

func (s *serviceService) GetById(ctx context.Context, id int) (*models.Service, error) {
	if id < 1 {
		return nil, errors.New("Некорректное значение id")
	}

	service := models.Service{}
	if err := s.serviceRepository.GetById(ctx, id, &service); err != nil {
		return nil, err
	}

	return &service, nil
}

func (s *serviceService) Update(c context.Context, id int, name string) (*models.Service, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	service := &models.Service{Id: id, Name: name}

	if err := s.serviceRepository.Update(ctx, service); err != nil {
		return nil, err
	}

	return service, nil
}

func (s *serviceService) Delete(c context.Context, id int) error {
	// TODO delete service service
	return nil
}
