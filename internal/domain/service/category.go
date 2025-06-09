package service

import (
	"context"
	"errors"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type categoryService struct {
	svc contract.ServiceManager
}

func NewCategoryService(svc contract.ServiceManager) contract.CategoryService {
	return &categoryService{
		svc: svc,
	}
}

func (s *categoryService) Save(ctx context.Context, request request.NewCategoryRequest) (categoryResponse response.CategoryResponse, err error) {
	if request.Name == "" {
		return response.CategoryResponse{}, errors.New("invalid category")
	}

	category := request.ToEntity()

	newCategory, err := s.svc.DB().CategoryRepo().Save(ctx, category)
	if err != nil {
		return response.CategoryResponse{}, err
	}
	categoryResponse = response.FromCategoryModel(newCategory)
	return categoryResponse, nil
}
