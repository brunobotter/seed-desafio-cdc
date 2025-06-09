package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type BookModel struct {
	ID          uint
	Title       string
	Resume      string
	Summary     string
	Price       float64
	Page        int64
	ISBN        string
	PublishDate time.Time
	AuthorID    uint
	CategoryID  uint
	CreatedAt   time.Time
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

type BookByIdModel struct {
	ID           uint
	Title        string
	Resume       string
	Summary      string
	Price        float64
	Page         int64
	ISBN         string
	PublishDate  time.Time
	AuthorName   string
	CategoryName string
	CreatedAt    time.Time
}

type BookByAllModel struct {
	ID    uint
	Title string
}
