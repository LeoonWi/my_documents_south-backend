package interfaces

type TariffService interface{}

type RoleService interface{}

type ServiceService interface{}

type ServiceFacade interface {
	Tariff() TariffService
	Role() RoleService
	Service() ServiceService
}
