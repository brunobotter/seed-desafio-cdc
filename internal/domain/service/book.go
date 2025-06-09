package service

import (
	"context"
	"errors"

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
	if request.Title == "" {
		return response.BookResponse{}, errors.New("title is required")
	}
	if request.Resume == "" {
		return response.BookResponse{}, errors.New("resume is required")
	}
	if len(request.Resume) > 500 {
		return response.BookResponse{}, errors.New("resume must be at most 500 characters")
	}
	if request.Price < 20 {
		return response.BookResponse{}, errors.New("price must be at least 20")
	}
	if request.Page < 100 {
		return response.BookResponse{}, errors.New("page must be at least 100")
	}
	if request.ISBN == "" {
		return response.BookResponse{}, errors.New("isbn is required")
	}
	if request.PublishDate.IsZero() {
		return response.BookResponse{}, errors.New("publish_date is required")
	}
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
	bookResponse = response.FromBookByIdModel(bookData)
	return bookResponse, nil
}

func (s *bookService) GetAll(ctx context.Context) (bookResponse response.BookListResponse, err error) {
	bookData, err := s.svc.DB().BookRepo().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	bookResponse = response.FromListBookModel(bookData)
	return bookResponse, nil
}
