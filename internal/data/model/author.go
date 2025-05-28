package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type AuthorModel struct {
	ID          uint
	Name        string
	Email       string
	Description string
	CreatedAt   time.Time
}

func ToAuthorModel(e entity.Author) AuthorModel {
	return AuthorModel{
		Name:        e.Name,
		Email:       e.Email,
		Description: e.Description,
	}
}
