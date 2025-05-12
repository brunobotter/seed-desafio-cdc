package entity

import "time"

type Author struct {
	Name        string
	Email       string
	Description string
}

type AuthorBase struct {
	Name        string
	Email       string
	Description string
	CreateAt    time.Time
}
