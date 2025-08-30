package service

import (
	"my_documents_south_backend/internal/core/repositories"
)

type RequestService struct {
	repositories repositories.Postgres
}

func NewRequestService(repositories repositories.Postgres) *RequestService {
	return &RequestService{repositories: repositories}
}
