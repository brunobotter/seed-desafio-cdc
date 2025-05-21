package service

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type bookService struct {
	svc contract.ServiceManager
}

func NewBookService(svc contract.ServiceManager) contract.BookService {
	return &bookService{
		svc: svc,
	}
}

func (s *bookService) Save(ctx context.Context, request request.NewBookRequest, categoryId int64, authorId int64) (bookResponse response.BookResponse, err error) {
	book := request.ToEntity(categoryId, authorId)
	bookData, err := s.svc.DB().BookRepo().Save(ctx, book)
	if err != nil {
		return response.BookResponse{}, err
	}
	bookResponse = response.FromBookModel(bookData)
	return bookResponse, nil
}

func (s *bookService) GetById(ctx context.Context, bookId int64) (bookResponse response.BookResponse, err error) {
	bookData, err := s.svc.DB().BookRepo().GetById(ctx, bookId)
	if err != nil {
		return response.BookResponse{}, err
	}
	bookResponse = response.FromBookModel(bookData)
	return bookResponse, nil
}
