package postgres

import "github.com/jmoiron/sqlx"

type (
	ITariffRepository interface{}

	IRoleRepository interface{}

	ISvcRepository interface{}

	Repository struct {
		TariffRepository ITariffRepository
		RoleRepository   IRoleRepository
		SvcRepository    SvcRepository
	}
)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TariffRepository: TariffRepository{db: db},
		RoleRepository:   RoleRepository{db: db},
		SvcRepository:    SvcRepository{db: db},
	}
}
