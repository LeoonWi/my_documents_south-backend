package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type TariffService struct {
	repositories repositories.Repositories
}

func NewTariffService(repositories repositories.Repositories) *TariffService {
	return &TariffService{repositories: repositories}
}
