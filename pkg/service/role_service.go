package service

import (
	"my_documents_south_backend/pkg/repository/postgres"
)

type RoleService struct {
	repository *postgres.Repository
}
