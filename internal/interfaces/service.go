package interfaces

import "my_documents_south_backend/internal/model"

type EmployeeService interface{}

type EmployeeSpecsService interface{}

type RequestService interface{}

type RoleService interface {
	CreateRole(name string) (*model.Role, error)
}

type ServiceService interface {
	CreateService(name string) (*model.Service, error)
}

type TariffService interface {
	CreateTariff(name string) (*model.Tariff, error)
}

type UserService interface{}
