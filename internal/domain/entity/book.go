package entity

import "time"

type Book struct {
	Title       string
	Resume      string
	Summary     string
	Price       float64
	Page        int64
	ISBN        string
	PublishDate time.Time
	CategoryId  int64
	AuthorId    int64
}
