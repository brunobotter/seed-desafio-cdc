package mocks_repo

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type AuthorRepoMock struct {
	SaveFunc func(ctx context.Context, author entity.Author) (model.AuthorModel, error)
}

func (m *AuthorRepoMock) Save(ctx context.Context, author entity.Author) (model.AuthorModel, error) {
	return m.SaveFunc(ctx, author)
}

type BookRepoMock struct {
	SaveFunc    func(ctx context.Context, book entity.Book) (model.BookModel, error)
	GetByIdFunc func(ctx context.Context, bookId int64) (model.BookByIdModel, error)
	GetAllFunc  func(ctx context.Context) ([]model.BookByAllModel, error)
}

func (m *BookRepoMock) Save(ctx context.Context, book entity.Book) (model.BookModel, error) {
	return m.SaveFunc(ctx, book)
}

func (m *BookRepoMock) GetById(ctx context.Context, bookId int64) (model.BookByIdModel, error) {
	return m.GetByIdFunc(ctx, bookId)
}

func (m *BookRepoMock) GetAll(ctx context.Context) ([]model.BookByAllModel, error) {
	return m.GetAllFunc(ctx)
}

type CategoryRepoMock struct {
	SaveFunc func(ctx context.Context, category entity.Category) (model.CategoryModel, error)
}

func (m *CategoryRepoMock) Save(ctx context.Context, category entity.Category) (model.CategoryModel, error) {
	return m.SaveFunc(ctx, category)
}

type CountryRepoMock struct {
	SaveFunc   func(ctx context.Context, country entity.Country) (model.CountryModel, error)
	VerifyFunc func(ctx context.Context, country string, state *string) (model.CountryState, error)
}

func (m *CountryRepoMock) Save(ctx context.Context, country entity.Country) (model.CountryModel, error) {
	return m.SaveFunc(ctx, country)
}

func (m *CountryRepoMock) VerifyCountryState(ctx context.Context, country string, state *string) (model.CountryState, error) {
	return m.VerifyFunc(ctx, country, state)
}

type StateRepoMock struct {
	SaveFunc func(ctx context.Context, state entity.State) (model.StateModel, error)
}

func (m *StateRepoMock) Save(ctx context.Context, state entity.State) (model.StateModel, error) {
	return m.SaveFunc(ctx, state)
}

type CustomerRepoMock struct {
	SaveFunc    func(ctx context.Context, payment entity.Customer) (model.CustomerModel, error)
	GetByIdFunc func(ctx context.Context, customerId int64) (model.CustomerModel, error)
}

func (m *CustomerRepoMock) Save(ctx context.Context, payment entity.Customer) (model.CustomerModel, error) {
	return m.SaveFunc(ctx, payment)
}

func (m *CustomerRepoMock) GetById(ctx context.Context, customerId int64) (model.CustomerModel, error) {
	return m.GetByIdFunc(ctx, customerId)
}
