package facade

import (
	"my_documents_south_backend/internal/core/interfaces"
	"my_documents_south_backend/internal/domain/role"
	"my_documents_south_backend/internal/domain/service"
	"my_documents_south_backend/internal/domain/tariff"

	"github.com/jmoiron/sqlx"
)

type RepositoryFacade struct {
	tariffRepository  interfaces.TariffRepository
	roleRepository    interfaces.RoleRepository
	serviceRepository interfaces.ServiceRepository
}

func NewRepository(db *sqlx.DB) *RepositoryFacade {
	return &RepositoryFacade{
		tariffRepository:  tariff.Repository{Conn: db},
		roleRepository:    role.Repository{Conn: db},
		serviceRepository: service.Repository{Conn: db},
	}
}

func (f *RepositoryFacade) Tariff() interfaces.TariffRepository {
	return f.tariffRepository
}

func (f *RepositoryFacade) Role() interfaces.RoleRepository {
	return f.roleRepository
}

func (f *RepositoryFacade) Service() interfaces.ServiceRepository {
	return f.serviceRepository
}
