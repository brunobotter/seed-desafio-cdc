package request

import "github.com/brunobotter/casa-codigo/internal/domain/entity"

type NewAuthorRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Description string `json:"description" binding:"required,max=400"`
}

func (r NewAuthorRequest) ToEntity() entity.Author {
	return entity.Author{
		Name:        r.Name,
		Email:       r.Email,
		Description: r.Description,
	}
}
