package services

import (
	"context"
	"errors"
	"my_documents_south_backend/internal/models"
	"time"
)

type tariffService struct {
	tariffRepository models.TariffRepository
	contextTimeout   time.Duration
}

func NewTariffService(tariffRepository models.TariffRepository, contextTimeout time.Duration) models.TariffService {
	return &tariffService{tariffRepository: tariffRepository, contextTimeout: contextTimeout}
}

func (s *tariffService) Create(c context.Context, name string) (*models.Tariff, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	tariff := &models.Tariff{Name: name}
	err := s.tariffRepository.Create(ctx, tariff)
	if err != nil {
		return nil, err
	}
	return tariff, nil
}

func (s *tariffService) Get(c context.Context) *[]models.Tariff {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	var tariff []models.Tariff
	err := s.tariffRepository.Get(ctx, &tariff)
	if err != nil {
		return nil
	}
	return &tariff
}

func (s *tariffService) GetById(c context.Context, id int) (*models.Tariff, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	if id < 1 {
		return nil, errors.New("Некорректное значение id")
	}

	tariff := &models.Tariff{Id: id}
	err := s.tariffRepository.GetById(ctx, id, tariff)
	if err != nil {
		return nil, err
	}
	return tariff, nil
}

func (s *tariffService) Update(c context.Context, id int, name string) (*models.Tariff, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	tariff := &models.Tariff{Id: id, Name: name}

	if err := s.tariffRepository.Update(ctx, tariff); err != nil {
		return nil, err
	}

	return tariff, nil
}

func (s *tariffService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	if id <= 0 {
		return errors.New("invalid id")
	}

	return s.tariffRepository.Delete(ctx, id)
}
