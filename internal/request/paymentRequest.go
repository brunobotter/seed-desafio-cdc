package request

import "github.com/brunobotter/casa-codigo/internal/domain/entity"

type NewPaymentRequest struct {
	Email      string  `json:"email" validate:"required,email"`
	Name       string  `json:"name" validate:"required"`
	LastName   string  `json:"lastname" validate:"required"`
	Document   string  `json:"document" validate:"required,cpf_cnpj"`
	Address    string  `json:"address" validate:"required"`
	Complement string  `json:"complement" validate:"required"`
	City       string  `json:"city" validate:"required"`
	Country    string  `json:"country" validate:"required"`
	State      *string `json:"state"`
	Phone      string  `json:"phone" validate:"required"`
	CEP        string  `json:"cep" validate:"required"`
}

func (r NewPaymentRequest) ToEntity() entity.Payment {
	return entity.Payment{
		Email:      r.Email,
		Name:       r.Name,
		Lastname:   r.LastName,
		Document:   r.Document,
		Address:    r.Address,
		Complement: r.Complement,
		City:       r.City,
		Country:    r.Country,
		State:      r.State,
		Phone:      r.Phone,
		CEP:        r.CEP,
	}
}
