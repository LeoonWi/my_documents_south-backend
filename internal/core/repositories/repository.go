package repositories

import (
	"my_documents_south_backend/internal/interfaces"
	"my_documents_south_backend/internal/repository"

	"github.com/jmoiron/sqlx"
)

type (
	Repositories interface {
		Employee() interfaces.EmployeeService
		EmployeeSpecs() interfaces.EmployeeSpecsService
		Request() interfaces.RequestService
		Role() interfaces.RoleService
		Service() interfaces.ServiceService
		Tariff() interfaces.TariffRepository
		User() interfaces.UserRepository
	}
	RepositoriesImpl struct {
		employeeRepository      interfaces.EmployeeRepository
		employeeSpecsRepository interfaces.EmployeeSpecsRepository
		requestRepository       interfaces.RequestRepository
		roleRepository          interfaces.RoleRepository
		serviceRepository       interfaces.ServiceRepository
		tariffRepository        interfaces.TariffRepository
		userRepository          interfaces.UserRepository
	}
)

func NewRepositories(db *sqlx.DB) *RepositoriesImpl {
	return &RepositoriesImpl{
		employeeRepository:      repository.NewEmployeeRepository(db),
		employeeSpecsRepository: repository.NewEmployeeSpecsRepository(db),
		requestRepository:       repository.NewRequestRepository(db),
		roleRepository:          repository.NewRoleRepository(db),
		serviceRepository:       repository.NewServiceRepository(db),
		tariffRepository:        repository.NewTariffRepository(db),
		userRepository:          repository.NewUserRepository(db),
	}
}

func (r *RepositoriesImpl) Employee() interfaces.EmployeeService {
	return r.employeeRepository
}

func (r *RepositoriesImpl) EmployeeSpecs() interfaces.EmployeeSpecsService {
	return r.employeeSpecsRepository
}

func (r *RepositoriesImpl) Request() interfaces.RequestService {
	return r.requestRepository
}

func (r *RepositoriesImpl) Role() interfaces.RoleService {
	return r.roleRepository
}

func (r *RepositoriesImpl) Service() interfaces.ServiceService {
	return r.serviceRepository
}

func (r *RepositoriesImpl) Tariff() interfaces.TariffRepository {
	return r.tariffRepository
}

func (r *RepositoriesImpl) User() interfaces.UserRepository {
	return r.userRepository
}
