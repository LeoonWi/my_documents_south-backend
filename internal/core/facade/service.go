package facade

import (
	"my_documents_south_backend/internal/core/interfaces"
	"my_documents_south_backend/internal/domain/role"
	"my_documents_south_backend/internal/domain/service"
	"my_documents_south_backend/internal/domain/tariff"
)

type ServiceFacade struct {
	tariffService  interfaces.TariffRepository
	roleService    interfaces.RoleRepository
	serviceService interfaces.ServiceRepository
}

func NewService(
	repository *RepositoryFacade,
) *ServiceFacade {
	return &ServiceFacade{
		tariffService:  tariff.Service{Repository: repository},
		roleService:    role.Service{Repository: repository},
		serviceService: service.Service{Repository: repository},
	}
}

func (s *ServiceFacade) Tariff() interfaces.TariffService {
	return s.tariffService
}

func (s *ServiceFacade) Role() interfaces.RoleService {
	return s.roleService
}
func (s *ServiceFacade) Service() interfaces.ServiceService {
	return s.serviceService
}
