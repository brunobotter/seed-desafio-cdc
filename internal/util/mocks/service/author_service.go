package mocks_service

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type AuthorServiceMock struct {
	SaveFunc func(ctx context.Context, req request.NewAuthorRequest) (response.AuthorResponse, error)
}

func (m *AuthorServiceMock) Save(ctx context.Context, req request.NewAuthorRequest) (response.AuthorResponse, error) {
	return m.SaveFunc(ctx, req)
}
