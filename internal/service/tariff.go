package service

import (
	"my_documents_south_backend/internal/core/repositories"
	"my_documents_south_backend/internal/model"
)

type TariffService struct {
	repositories repositories.Postgres
}

func NewTariffService(repositories repositories.Postgres) *TariffService {
	return &TariffService{repositories: repositories}
}

func (s *TariffService) CreateTariff(name string) (*model.Tariff, error) {
	return nil, nil
}
