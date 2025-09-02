package interfaces

import (
	"context"
	"my_documents_south_backend/internal/model"
)

type EmployeeRepository interface{}

type EmployeeSpecsRepository interface{}

type RequestRepository interface{}

type RoleRepository interface {
	CreateRole(ctx context.Context, role *model.Role) error
	GetRoles(ctx context.Context, roles *[]model.Role) error
	GetRoleById(ctx context.Context, id int, role *model.Role) error
	GetRoleByName(ctx context.Context, name string, role *model.Role) error
}

type ServiceRepository interface {
	CreateService(ctx context.Context, service *model.Service) error
	GetService(ctx context.Context, service *[]model.Service) error
	GetServiceById(ctx context.Context, id int, service *model.Service) error
	GetServiceByName(ctx context.Context, name string, service *model.Service) error
}

type TariffRepository interface {
	CreateTariff(ctx context.Context, role *model.Tariff) error
	GetTariffById(ctx context.Context, id int, tariff *model.Tariff) error
	GetTariffByName(ctx context.Context, name string, tariff *model.Tariff) error
}

type UserRepository interface{}
