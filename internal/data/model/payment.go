package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type PaymentModel struct {
	ID         uint
	Email      string
	Name       string
	Lastname   string
	Document   string
	Address    string
	Complement string
	City       string
	Country    string
	State      *string
	Phone      string
	CEP        string
	CreatedAt  time.Time
}

func ToPaymentModel(e entity.Payment) PaymentModel {
	return PaymentModel{
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
