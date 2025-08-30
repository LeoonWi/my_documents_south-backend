package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type RequestService struct {
	repositories repositories.Repositories
}

func NewRequestService(repositories repositories.Repositories) *RequestService {
	return &RequestService{repositories: repositories}
}
