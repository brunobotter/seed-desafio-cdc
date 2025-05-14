package request

import "github.com/brunobotter/casa-codigo/internal/domain/entity"

type NewCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (r NewCategoryRequest) ToEntity() entity.Category {
	return entity.Category{
		Name: r.Name,
	}
}
