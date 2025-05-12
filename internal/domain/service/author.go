package service

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
)

type authorInternalService struct {
	svc contract.ServiceManager
}

func NewAuthorIntegrationService(svc contract.ServiceManager) contract.AuthorService {
	return &authorInternalService{
		svc: svc,
	}
}

func (s *authorInternalService) Save(ctx context.Context, request request.NewAuthorRequest) error {
	return nil
}
