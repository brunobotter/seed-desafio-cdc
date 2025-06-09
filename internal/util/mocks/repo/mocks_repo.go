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
