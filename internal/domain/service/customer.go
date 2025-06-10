package service

import (
	"context"
	"errors"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
	"github.com/go-playground/validator/v10"
)

type customerService struct {
	svc contract.ServiceManager
}

var validate = validator.New()

func NewCustomerService(svc contract.ServiceManager) contract.CustomerService {
	return &customerService{
		svc: svc,
	}
}

func (s *customerService) Save(ctx context.Context, request request.NewCustomerRequest) (customerResponse response.CustomerResponse, err error) {
	if err := validate.Struct(request); err != nil {
		return response.CustomerResponse{}, err
	}
	if request.Countryid == "" {
		return response.CustomerResponse{}, errors.New("invalid country")
	}
	data, err := s.svc.DB().CountryRepo().VerifyCountryState(ctx, request.Countryid, request.Stateid)
	if err != nil {
		return response.CustomerResponse{}, err
	}

	customer := request.ToEntity(data)

	customerData, err := s.svc.DB().CustomerRepo().Save(ctx, customer)
	if err != nil {
		return response.CustomerResponse{}, err
	}
	customerResponse = response.FromCustomerModel(customerData)
	return customerResponse, nil
}
