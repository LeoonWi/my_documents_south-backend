package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type TariffService struct {
	repositories repositories.Postgres
}

func NewTariffService(repositories repositories.Postgres) *TariffService {
	return &TariffService{repositories: repositories}
}
