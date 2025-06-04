package response

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
)

type CustomerResponse struct {
	ID         uint      `json:"id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Lastname   string    `json:"lastname"`
	Document   string    `json:"document"`
	Address    string    `json:"address"`
	Complement string    `json:"complement"`
	City       string    `json:"city"`
	Country    int64     `json:"country"`
	State      *int64    `json:"state,omitempty"`
	Phone      string    `json:"phone"`
	CEP        string    `json:"cep"`
	CreatedAt  time.Time `json:"created_at"`
}

func FromCustomerModel(m model.CustomerModel) CustomerResponse {
	return CustomerResponse{
		ID:         m.ID,
		Email:      m.Email,
		Name:       m.Name,
		Lastname:   m.Lastname,
		Document:   m.Document,
		Address:    m.Address,
		Complement: m.Complement,
		City:       m.City,
		Country:    m.Country,
		State:      m.State,
		Phone:      m.Phone,
		CEP:        m.CEP,
		CreatedAt:  m.CreatedAt,
	}
}
