package service

import (
	"context"
	"errors"
	"regexp"

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
	if request.Name == "" {
		return response.AuthorResponse{}, errors.New("invalid name")
	}
	if request.Description == "" || len(request.Description) > 400 {
		return response.AuthorResponse{}, errors.New("invalid description")
	}
	if !isValidEmail(request.Email) {
		return response.AuthorResponse{}, errors.New("invalid email")
	}
	authorDb, err := s.svc.DB().AuthorRepo().Save(ctx, author)
	if err != nil {
		return response.AuthorResponse{}, err
	}

	authorResponse = response.FromAuthorModel(authorDb)
	return authorResponse, nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
