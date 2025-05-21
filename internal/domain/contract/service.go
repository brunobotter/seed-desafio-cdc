package contract

import (
	"context"

	"github.com/brunobotter/casa-codigo/configs/mapping"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type ServiceManager interface {
	Config() *mapping.Config
	DB() DataManager
	InternalService() InternalService
}

type InternalService interface {
	AuthorService() AuthorService
	CategoryService() CategoryService
	BookService() BookService
}

type AuthorService interface {
	Save(ctx context.Context, request request.NewAuthorRequest) (response response.AuthorResponse, err error)
}

type CategoryService interface {
	Save(ctx context.Context, request request.NewCategoryRequest) (response response.CategoryResponse, err error)
}

type BookService interface {
	Save(ctx context.Context, request request.NewBookRequest, categoryId int64, authorId int64) (response response.BookResponse, err error)
}
