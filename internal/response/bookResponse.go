package response

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
)

type BookResponse struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Resume       string    `json:"resume"`
	Summary      string    `json:"summary"`
	Price        float64   `json:"price"`
	Page         int64     `json:"page"`
	ISBN         string    `json:"isbn"`
	PublishDate  time.Time `json:"publish_date"`
	AuthorName   string    `json:"author_name"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
}

type BookListResponse []BookItemResponse

type BookItemResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func FromBookModel(m model.BookModel) BookResponse {
	return BookResponse{
		ID:          m.ID,
		Title:       m.Title,
		Resume:      m.Resume,
		Summary:     m.Summary,
		Price:       m.Price,
		Page:        m.Page,
		ISBN:        m.ISBN,
		PublishDate: m.PublishDate,
		CreatedAt:   m.CreatedAt,
	}
}

func FromBookByIdModel(m model.BookByIdModel) BookResponse {
	return BookResponse{
		ID:           m.ID,
		Title:        m.Title,
		Resume:       m.Resume,
		Summary:      m.Summary,
		Price:        m.Price,
		Page:         m.Page,
		ISBN:         m.ISBN,
		PublishDate:  m.PublishDate,
		AuthorName:   m.AuthorName,
		CategoryName: m.CategoryName,
		CreatedAt:    m.CreatedAt,
	}
}

func FromListBookModel(models []model.BookByAllModel) BookListResponse {
	var response BookListResponse

	for _, m := range models {
		response = append(response, BookItemResponse{
			ID:    m.ID,
			Title: m.Title,
		})
	}

	return response
}
