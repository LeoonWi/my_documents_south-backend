package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type UserService struct {
	repositories repositories.Postgres
}

func NewUserService(repositories repositories.Postgres) *UserService {
	return &UserService{repositories: repositories}
}
