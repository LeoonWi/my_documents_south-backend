package service

import "my_documents_south_backend/pkg/repository/postgres"

type SvcService struct {
	repository *postgres.Repository
}
