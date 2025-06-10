package mocks_service

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
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

type CountryServiceMock struct {
	SaveFunc func(ctx context.Context, request request.NewCountryRequest) (response response.CountryResponse, err error)
}

func (m *CountryServiceMock) Save(ctx context.Context, request request.NewCountryRequest) (response response.CountryResponse, err error) {
	return m.SaveFunc(ctx, request)
}

type StateServiceMock struct {
	SaveFunc func(ctx context.Context, state entity.State) (model.StateModel, error)
}

func (m *StateServiceMock) Save(ctx context.Context, state entity.State) (model.StateModel, error) {
	return m.SaveFunc(ctx, state)
}

type CustomerServiceMock struct {
	SaveFunc func(ctx context.Context, request request.NewCustomerRequest) (response response.CustomerResponse, err error)
}

func (m *CustomerServiceMock) Save(ctx context.Context, request request.NewCustomerRequest) (response response.CustomerResponse, err error) {
	return m.SaveFunc(ctx, request)
}
