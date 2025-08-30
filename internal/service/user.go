package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type UserService struct {
	repositories repositories.Repositories
}

func NewUserService(repositories repositories.Repositories) *UserService {
	return &UserService{repositories: repositories}
}
