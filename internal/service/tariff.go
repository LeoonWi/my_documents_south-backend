package service

import (
	"context"
	"errors"
	"my_documents_south_backend/internal/core/repositories"
	"my_documents_south_backend/internal/model"
)

type TariffService struct {
	repositories repositories.Postgres
}

func NewTariffService(repositories repositories.Postgres) *TariffService {
	return &TariffService{repositories: repositories}
}

func (s *TariffService) CreateTariff(ctx context.Context, name string) (*model.Tariff, error) {
	tariff := &model.Tariff{Name: name}
	err := s.repositories.Tariff().CreateTariff(ctx, tariff)
	if err != nil {
		return nil, err
	}
	return tariff, nil
}

func (s *TariffService) GetTariffs(ctx context.Context) *[]model.Tariff {
	var tariff []model.Tariff
	err := s.repositories.Tariff().GetTariffs(ctx, &tariff)
	if err != nil {
		return nil
	}
	return nil
}

func (s *TariffService) GetTariffByID(ctx context.Context, id int) (*model.Tariff, error) {
	if id < 1 {
		return nil, errors.New("Некорректное значение id")
	}

	tariff := &model.Tariff{Id: id}
	err := s.repositories.Tariff().GetTariffByID(ctx, id, tariff)
	if err != nil {
		return nil, err
	}
	return tariff, nil
}
func (s *TariffService) GetTariffByName(ctx context.Context, name string) (*model.Tariff, error) {
	tariff := &model.Tariff{Name: name}
	err := s.repositories.Tariff().GetTariffByName(ctx, name, tariff)
	if err != nil {
		return nil, err
	}
	return tariff, nil
}

func (s *TariffService) UpdateTariff(ctx context.Context, id int, name string) (*model.Tariff, error) {
	tariff := &model.Tariff{Id: id, Name: name}

	if err := s.repositories.Tariff().UpdateTariff(ctx, tariff); err != nil {
		return nil, err
	}

	return tariff, nil
}

func (s *TariffService) DeleteTariff(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	return s.repositories.Tariff().DeleteTariff(ctx, id)
}
