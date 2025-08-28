package service

import "my_documents_south_backend/pkg/repository/postgres"

type (
	ITariffService interface{}

	IRoleService interface{}

	ISvcService interface{}

	Service struct {
		TariffService ITariffService
		RoleService   IRoleService
		SvcService    ISvcService
	}
)

func NewService(
	repository *postgres.Repository,
) *Service {
	return &Service{
		TariffService: TariffService{repository: repository},
		RoleService:   RoleService{repository: repository},
		SvcService:    SvcService{repository: repository},
	}
}
