package services

import (
	"my_documents_south_backend/internal/core/repositories"
	"my_documents_south_backend/internal/interfaces"
	"my_documents_south_backend/internal/service"
)

type (
	Service interface {
		Employee() interfaces.EmployeeService
		EmployeeSpecs() interfaces.EmployeeSpecsService
		Request() interfaces.RequestService
		Role() interfaces.RoleService
		Service() interfaces.ServiceService
		Tariff() interfaces.TariffService
		User() interfaces.UserService
	}
	ServiceImpl struct {
		employeeService      interfaces.EmployeeService
		employeeSpecsService interfaces.EmployeeSpecsService
		requestService       interfaces.RequestService
		roleService          interfaces.RoleService
		serviceService       interfaces.ServiceService
		tariffService        interfaces.TariffService
		userService          interfaces.UserService
	}
)

func NewService(repositories repositories.Postgres) *ServiceImpl {
	return &ServiceImpl{
		employeeService:      service.NewEmployeeService(repositories),
		employeeSpecsService: service.NewEmployeeSpecsService(repositories),
		requestService:       service.NewRequestService(repositories),
		roleService:          service.NewRoleService(repositories),
		serviceService:       service.NewServiceService(repositories),
		tariffService:        service.NewTariffService(repositories),
		userService:          service.NewUserService(repositories),
	}
}

func (s *ServiceImpl) Employee() interfaces.EmployeeService {
	return s.employeeService
}

func (s *ServiceImpl) EmployeeSpecs() interfaces.EmployeeSpecsService {
	return s.employeeService
}

func (s *ServiceImpl) Request() interfaces.RequestService {
	return s.requestService
}

func (s *ServiceImpl) Role() interfaces.RoleService {
	return s.roleService
}

func (s *ServiceImpl) Service() interfaces.ServiceService {
	return s.serviceService
}

func (s *ServiceImpl) Tariff() interfaces.TariffService {
	return s.tariffService
}

func (s *ServiceImpl) User() interfaces.UserService {
	return s.userService
}
