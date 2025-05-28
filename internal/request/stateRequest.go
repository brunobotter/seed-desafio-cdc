package request

import "github.com/brunobotter/casa-codigo/internal/domain/entity"

type NewStateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (r NewStateRequest) ToEntity(countryId int64) entity.State {
	return entity.State{
		Name:      r.Name,
		CountryId: countryId,
	}
}
