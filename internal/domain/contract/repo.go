package contract

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type RepoManager interface {
	AuthorRepo() AuthorRepository
	CategoryRepo() CategoryRepository
	BookRepo() BookRepository
}

type AuthorRepository interface {
	Save(ctx context.Context, author entity.Author) (model.AuthorModel, error)
}

type CategoryRepository interface {
	Save(ctx context.Context, category entity.Category) (model.CategoryModel, error)
}

type BookRepository interface {
	Save(ctx context.Context, book entity.Book) (model.BookModel, error)
	GetById(ctx context.Context, bookId int64) (model.BookModel, error)
	GetAll(ctx context.Context) ([]model.BookModel, error)
}
