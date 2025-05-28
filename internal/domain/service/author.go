package service

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type authorService struct {
	svc contract.ServiceManager
}

func NewAuthorService(svc contract.ServiceManager) contract.AuthorService {
	return &authorService{
		svc: svc,
	}
}

func (s *authorService) Save(ctx context.Context, request request.NewAuthorRequest) (authorResponse response.AuthorResponse, err error) {
	author := request.ToEntity()
	authorDb, err := s.svc.DB().AuthorRepo().Save(ctx, author)
	if err != nil {
		return response.AuthorResponse{}, err
	}

	authorResponse = response.FromAuthorModel(authorDb)
	return authorResponse, nil
}
