package response

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
)

type CountryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func FromCountryModel(m model.CountryModel) CountryResponse {
	return CountryResponse{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}
