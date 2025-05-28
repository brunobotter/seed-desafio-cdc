package service

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type countryService struct {
	svc contract.ServiceManager
}

func NewCountryService(svc contract.ServiceManager) contract.CountryService {
	return &countryService{
		svc: svc,
	}
}

func (s *countryService) Save(ctx context.Context, request request.NewCountryRequest) (countryResponse response.CountryResponse, err error) {
	country := request.ToEntity()
	countryDb, err := s.svc.DB().CountryRepo().Save(ctx, country)
	if err != nil {
		return response.CountryResponse{}, err
	}

	countryResponse = response.FromCountryModel(countryDb)
	return countryResponse, nil
}
