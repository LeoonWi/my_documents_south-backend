package interfaces

type TariffRepository interface{}

type RoleRepository interface{}

type ServiceRepository interface{}

type RepositoryFacade interface {
	Tariff() TariffRepository
	Role() RoleRepository
	Service() ServiceRepository
}
