package service

import (
	"context"
	"errors"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type stateService struct {
	svc contract.ServiceManager
}

func NewStateService(svc contract.ServiceManager) contract.StateService {
	return &stateService{
		svc: svc,
	}
}

func (s *stateService) Save(ctx context.Context, request request.NewStateRequest, countryId int64) (stateResponse response.StateResponse, err error) {
	if request.Name == "" {
		return response.StateResponse{}, errors.New("invalid state")
	}
	state := request.ToEntity(countryId)
	stateDb, err := s.svc.DB().StateRepo().Save(ctx, state)
	if err != nil {
		return response.StateResponse{}, err
	}

	stateResponse = response.FromStateModel(stateDb)
	return stateResponse, nil
}
