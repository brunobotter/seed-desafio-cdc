package request

import "github.com/brunobotter/casa-codigo/internal/domain/entity"

type NewCountryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (r NewCountryRequest) ToEntity() entity.Country {
	return entity.Country{
		Name: r.Name,
	}
}
