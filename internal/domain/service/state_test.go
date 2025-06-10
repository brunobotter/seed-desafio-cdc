package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
	"github.com/brunobotter/casa-codigo/internal/domain/service"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util/mocks"
	mocks_repo "github.com/brunobotter/casa-codigo/internal/util/mocks/repo"
	"github.com/stretchr/testify/require"
)

func TestStateService_Save(t *testing.T) {
	tests := []struct {
		name          string
		countryId     int64
		req           request.NewStateRequest
		setupRepoMock func() *mocks_repo.StateRepoMock
		wantErr       bool
		wantErrMsg    string
	}{
		{
			name:          "Nome vazio",
			countryId:     1,
			req:           request.NewStateRequest{Name: ""},
			setupRepoMock: func() *mocks_repo.StateRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "invalid state",
		},
		{
			name:      "Repo retorna erro (decisão do repo)",
			countryId: 1,
			req:       request.NewStateRequest{Name: "test"},
			setupRepoMock: func() *mocks_repo.StateRepoMock {
				return &mocks_repo.StateRepoMock{
					SaveFunc: func(ctx context.Context, state entity.State) (model.StateModel, error) {
						return model.StateModel{}, errors.New("db error")
					},
				}
			},
			wantErr:    true,
			wantErrMsg: "db error",
		},
		{
			name:      "Sucesso",
			countryId: 1,
			req:       request.NewStateRequest{Name: "Fulano"},
			setupRepoMock: func() *mocks_repo.StateRepoMock {
				return &mocks_repo.StateRepoMock{
					SaveFunc: func(ctx context.Context, state entity.State) (model.StateModel, error) {
						return model.StateModel{
							ID:        1,
							Statename: state.Name,
							CountryId: 1,
						}, nil
					},
				}
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var repoMock contract.StateRepository
			if tc.setupRepoMock != nil {
				repoMock = tc.setupRepoMock()
			}
			mockDataManager := &mocks.DataManagerMock{
				StateRepoField: repoMock,
			}
			mockServiceManager := &mocks.ServiceManagerMock{
				DataManagerField: mockDataManager,
			}
			svc := service.NewStateService(mockServiceManager)
			resp, err := svc.Save(context.Background(), tc.req, tc.countryId)

			if tc.wantErr {
				require.Error(t, err)
				if tc.wantErrMsg != "" {
					require.Contains(t, err.Error(), tc.wantErrMsg)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.req.Name, resp.StateName)
				require.Equal(t, tc.countryId, resp.CountryId)
			}
		})
	}
}

/*
func TestCustomerService_Save_MissingRequiredFields(t *testing.T) {
    ctx := context.Background()
    svc := new(mockServiceManager)
    db := new(mockDB)
    countryRepo := new(mockCountryRepo)
    customerRepo := new(mockCustomerRepo)
    db.On("CountryRepo").Return(countryRepo)
    db.On("CustomerRepo").Return(customerRepo)
    svc.On("DB").Return(db)

    // Exemplo: Email vazio
    req := request.NewCustomerRequest{
        Email:      "",
        Name:       "John",
        LastName:   "Doe",
        Document:   "12345678901",
        Address:    "Rua 1",
        Complement: "Apto 1",
        City:       "Cidade",
        Countryid:  "1",
        Stateid:    nil,
        Phone:      "11999999999",
        CEP:        "12345678",
    }

    service := service.NewCustomerService(svc)
    resp, err := service.Save(ctx, req)
    assert.Error(t, err)
    assert.Equal(t, response.CustomerResponse{}, resp)
}

func TestCustomerService_Save_InvalidEmailFormat(t *testing.T) {
    ctx := context.Background()
    svc := new(mockServiceManager)
    db := new(mockDB)
    countryRepo := new(mockCountryRepo)
    customerRepo := new(mockCustomerRepo)
    db.On("CountryRepo").Return(countryRepo)
    db.On("CustomerRepo").Return(customerRepo)
    svc.On("DB").Return(db)

    req := request.NewCustomerRequest{
        Email:      "invalid-email",
        Name:       "John",
        LastName:   "Doe",
        Document:   "12345678901",
        Address:    "Rua 1",
        Complement: "Apto 1",
        City:       "Cidade",
        Countryid:  "1",
        Stateid:    nil,
        Phone:      "11999999999",
        CEP:        "12345678",
    }

    service := service.NewCustomerService(svc)
    resp, err := service.Save(ctx, req)
    assert.Error(t, err)
    assert.Equal(t, response.CustomerResponse{}, resp)
}

func TestCustomerService_Save_InvalidDocument(t *testing.T) {
    ctx := context.Background()
    svc := new(mockServiceManager)
    db := new(mockDB)
    countryRepo := new(mockCountryRepo)
    customerRepo := new(mockCustomerRepo)
    db.On("CountryRepo").Return(countryRepo)
    db.On("CustomerRepo").Return(customerRepo)
    svc.On("DB").Return(db)

    req := request.NewCustomerRequest{
        Email:      "test@email.com",
        Name:       "John",
        LastName:   "Doe",
        Document:   "invalid-doc",
        Address:    "Rua 1",
        Complement: "Apto 1",
        City:       "Cidade",
        Countryid:  "1",
        Stateid:    nil,
        Phone:      "11999999999",
        CEP:        "12345678",
    }

    service := service.NewCustomerService(svc)
    resp, err := service.Save(ctx, req)
    assert.Error(t, err)
    assert.Equal(t, response.CustomerResponse{}, resp)
}

func TestCustomerService_Save_StateRequiredWhenCountryHasStates(t *testing.T) {
    ctx := context.Background()
    svc := new(mockServiceManager)
    db := new(mockDB)
    countryRepo := new(mockCountryRepo)
    customerRepo := new(mockCustomerRepo)
    db.On("CountryRepo").Return(countryRepo)
    db.On("CustomerRepo").Return(customerRepo)
    svc.On("DB").Return(db)

    req := request.NewCustomerRequest{
        Email:      "test@email.com",
        Name:       "John",
        LastName:   "Doe",
        Document:   "12345678901",
        Address:    "Rua 1",
        Complement: "Apto 1",
        City:       "Cidade",
        Countryid:  "1",
        Stateid:    nil, // Estado não informado
        Phone:      "11999999999",
        CEP:        "12345678",
    }

    // Simula que o país tem estados e o estado não foi informado
    countryRepo.On("VerifyCountryState", ctx, "1", (*string)(nil)).Return(nil, errors.New("state required for this country"))

    service := service.NewCustomerService(svc)
    resp, err := service.Save(ctx, req)
    assert.Error(t, err)
    assert.Equal(t, response.CustomerResponse{}, resp)
}

func TestCustomerService_Save_CustomerRepoSaveError(t *testing.T) {
    ctx := context.Background()
    svc := new(mockServiceManager)
    db := new(mockDB)
    countryRepo := new(mockCountryRepo)
    customerRepo := new(mockCustomerRepo)
    db.On("CountryRepo").Return(countryRepo)
    db.On("CustomerRepo").Return(customerRepo)
    svc.On("DB").Return(db)

    req := request.NewCustomerRequest{
        Email:      "test@email.com",
        Name:       "John",
        LastName:   "Doe",
        Document:   "12345678901",
        Address:    "Rua 1",
        Complement: "Apto 1",
        City:       "Cidade",
        Countryid:  "1",
        Stateid:    nil,
        Phone:      "11999999999",
        CEP:        "12345678",
    }

    countryRepo.On("VerifyCountryState", ctx, "1", (*string)(nil)).Return("country-data", nil)
    customerRepo.On("Save", ctx, mock.Anything).Return(nil, errors.New("db error"))

    service := service.NewCustomerService(svc)
    resp, err := service.Save(ctx, req)
    assert.Error(t, err)
    assert.Equal(t, response.CustomerResponse{}, resp)
}*/
