package request

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type NewBookRequest struct {
	Title       string    `json:"title" validate:"required"`
	Resume      string    `json:"resume" validate:"required,max=500"`
	Summary     string    `json:"summary"`
	Price       float64   `json:"price" validate:"required,gte=20"`
	Page        int64     `json:"page" validate:"required,gte=100"`
	ISBN        string    `json:"isbn" validate:"required"`
	PublishDate time.Time `json:"publish_date" binding:"required" time_format:"2006-01-02"`
}

func (r NewBookRequest) ToEntity(categoryId int64, authorId int64) entity.Book {
	return entity.Book{
		Title:       r.Title,
		Resume:      r.Resume,
		Summary:     r.Summary,
		Price:       r.Price,
		Page:        r.Page,
		ISBN:        r.ISBN,
		PublishDate: r.PublishDate.Format("2006-01-02"),
		CategoryId:  categoryId,
		AuthorId:    authorId,
	}
}
