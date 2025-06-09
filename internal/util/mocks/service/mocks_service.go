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

type BookServiceMock struct {
	SaveFunc    func(ctx context.Context, req request.NewBookRequest) (response.BookResponse, error)
	GetByIdFunc func(ctx context.Context, bookId int64) (bookResponse response.BookResponse, err error)
	GetAllFunc  func(ctx context.Context) (bookResponse response.BookListResponse, err error)
}

func (m *BookServiceMock) Save(ctx context.Context, req request.NewBookRequest) (response.BookResponse, error) {
	return m.SaveFunc(ctx, req)
}

func (m *BookServiceMock) GetById(ctx context.Context, bookId int64) (bookResponse response.BookResponse, err error) {
	return m.GetByIdFunc(ctx, bookId)
}

func (m *BookServiceMock) GetAll(ctx context.Context) (bookResponse response.BookListResponse, err error) {
	return m.GetAllFunc(ctx)
}
