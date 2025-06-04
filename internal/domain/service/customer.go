package service

import (
	"context"
	"fmt"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type customerService struct {
	svc contract.ServiceManager
}

func NewCustomerService(svc contract.ServiceManager) contract.CustomerService {
	return &customerService{
		svc: svc,
	}
}

func (s *customerService) Save(ctx context.Context, request request.NewCustomerRequest) (customerResponse response.CustomerResponse, err error) {
	data, err := s.svc.DB().CountryRepo().VerifyCountryState(ctx, request.Countryid, request.Stateid)
	if err != nil {
		return response.CustomerResponse{}, err
	}
	customer := request.ToEntity(data)

	customerData, err := s.svc.DB().CustomerRepo().Save(ctx, customer)
	if err != nil {
		return response.CustomerResponse{}, err
	}
	fmt.Printf("test %v", customerData)
	customerResponse = response.FromCustomerModel(customerData)
	return customerResponse, nil
}
