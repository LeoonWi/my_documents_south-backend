package repositories

import (
	"my_documents_south_backend/internal/interfaces"
	"my_documents_south_backend/internal/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type (
	Postgres interface {
		Employee() interfaces.EmployeeService
		EmployeeSpecs() interfaces.EmployeeSpecsService
		Request() interfaces.RequestService
		Role() interfaces.RoleService
		Service() interfaces.ServiceService
		Tariff() interfaces.TariffRepository
		User() interfaces.UserRepository
	}
	PostgresImpl struct {
		employeeRepository      interfaces.EmployeeRepository
		employeeSpecsRepository interfaces.EmployeeSpecsRepository
		requestRepository       interfaces.RequestRepository
		roleRepository          interfaces.RoleRepository
		serviceRepository       interfaces.ServiceRepository
		tariffRepository        interfaces.TariffRepository
		userRepository          interfaces.UserRepository
	}
)

func NewRepositories(db *sqlx.DB) *PostgresImpl {
	return &PostgresImpl{
		employeeRepository:      postgres.NewEmployeeRepository(db),
		employeeSpecsRepository: postgres.NewEmployeeSpecsRepository(db),
		requestRepository:       postgres.NewRequestRepository(db),
		roleRepository:          postgres.NewRoleRepository(db),
		serviceRepository:       postgres.NewServiceRepository(db),
		tariffRepository:        postgres.NewTariffRepository(db),
		userRepository:          postgres.NewUserRepository(db),
	}
}

func (r *PostgresImpl) Employee() interfaces.EmployeeService {
	return r.employeeRepository
}

func (r *PostgresImpl) EmployeeSpecs() interfaces.EmployeeSpecsService {
	return r.employeeSpecsRepository
}

func (r *PostgresImpl) Request() interfaces.RequestService {
	return r.requestRepository
}

func (r *PostgresImpl) Role() interfaces.RoleService {
	return r.roleRepository
}

func (r *PostgresImpl) Service() interfaces.ServiceService {
	return r.serviceRepository
}

func (r *PostgresImpl) Tariff() interfaces.TariffRepository {
	return r.tariffRepository
}

func (r *PostgresImpl) User() interfaces.UserRepository {
	return r.userRepository
}
