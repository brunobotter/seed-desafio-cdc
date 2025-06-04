package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type CustomerModel struct {
	ID         uint
	Email      string
	Name       string
	Lastname   string
	Document   string
	Address    string
	Complement string
	City       string
	Country    int64
	State      *int64
	Phone      string
	CEP        string
	CreatedAt  time.Time
}

func ToCustomerModel(e entity.Customer) CustomerModel {
	return CustomerModel{
		Email:      e.Email,
		Name:       e.Name,
		Lastname:   e.Lastname,
		Document:   e.Document,
		Address:    e.Address,
		Complement: e.Complement,
		City:       e.City,
		Country:    e.Country,
		State:      e.State,
		Phone:      e.Phone,
		CEP:        e.CEP,
	}
}
