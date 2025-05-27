package response

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
)

type CategoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func FromCategoryModel(m model.CategoryModel) CategoryResponse {
	return CategoryResponse{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}
