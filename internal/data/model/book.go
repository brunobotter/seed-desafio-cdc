package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type BookModel struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Resume      string
	Summary     string
	Price       float64
	Page        int64
	ISBN        string
	PublishDate string
	CreatedAt   time.Time

	AuthorID uint
	Author   AuthorModel `gorm:"foreignKey:AuthorID"`

	CategoryID uint
	Category   CategoryModel `gorm:"foreignKey:CategoryID"`
}

func ToBookModel(e entity.Book) BookModel {
	return BookModel{
		Title:       e.Title,
		Resume:      e.Resume,
		Summary:     e.Summary,
		Price:       e.Price,
		Page:        e.Page,
		ISBN:        e.ISBN,
		PublishDate: e.PublishDate,
		AuthorID:    uint(e.AuthorId),
		CategoryID:  uint(e.CategoryId),
	}
}
