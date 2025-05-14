package response

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
)

type AuthorResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func FromAuthorModel(m model.AuthorModel) AuthorResponse {
	return AuthorResponse{
		ID:          m.AuthorID,
		Name:        m.Name,
		Email:       m.Email,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
	}
}
