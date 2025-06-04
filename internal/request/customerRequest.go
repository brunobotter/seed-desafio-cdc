package request

import (
	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type NewCustomerRequest struct {
	Email      string  `json:"email" validate:"required,email"`
	Name       string  `json:"name" validate:"required"`
	LastName   string  `json:"lastname" validate:"required"`
	Document   string  `json:"document" validate:"required,cpf_cnpj"`
	Address    string  `json:"address" validate:"required"`
	Complement string  `json:"complement" validate:"required"`
	City       string  `json:"city" validate:"required"`
	Countryid  string  `json:"country" validate:"required"`
	Stateid    *string `json:"state"`
	Phone      string  `json:"phone" validate:"required"`
	CEP        string  `json:"cep" validate:"required"`
}

func (r NewCustomerRequest) ToEntity(data model.CountryState) entity.Customer {
	return entity.Customer{
		Email:      r.Email,
		Name:       r.Name,
		Lastname:   r.LastName,
		Document:   r.Document,
		Address:    r.Address,
		Complement: r.Complement,
		City:       r.City,
		Country:    data.Countryid,
		State:      data.Stateid,
		Phone:      r.Phone,
		CEP:        r.CEP,
	}
}
