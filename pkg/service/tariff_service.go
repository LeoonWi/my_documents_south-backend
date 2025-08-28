package service

import "my_documents_south_backend/pkg/repository/postgres"

type TariffService struct {
	repository *postgres.Repository
}
