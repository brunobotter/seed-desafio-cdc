package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type AuthorModel struct {
	AuthorID    uint `gorm:"primaryKey"`
	Name        string
	Email       string `gorm:"unique"`
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
