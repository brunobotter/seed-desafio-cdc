package service_test

import (
	"context"
	"testing"

	"github.com/brunobotter/casa-codigo/internal/domain/service"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util/mocks"
	mocks_repo "github.com/brunobotter/casa-codigo/internal/util/mocks/repo"
	"github.com/stretchr/testify/require"
)

func TestCustomerService_Save(t *testing.T) {
	tests := []struct {
		name       string
		req        request.NewCustomerRequest
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "Campos obrigatórios faltando (email vazio)",
			req: request.NewCustomerRequest{
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
			},
			wantErr:    true,
			wantErrMsg: "Email",
		},
		{
			name: "País obrigatório",
			req: request.NewCustomerRequest{
				Email:      "test@email.com",
				Name:       "John",
				LastName:   "Doe",
				Document:   "12345678901",
				Address:    "Rua 1",
				Complement: "Apto 1",
				City:       "Cidade",
				Countryid:  "",
				Stateid:    nil,
				Phone:      "11999999999",
				CEP:        "12345678",
			},
			wantErr:    true,
			wantErrMsg: "invalid country",
		},
		{
			name: "Erro no VerifyCountryState",
			req: request.NewCustomerRequest{
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
			},
			wantErr:    true,
			wantErrMsg: "country/state error",
		},
		{
			name: "Erro ao salvar no CustomerRepo",
			req: request.NewCustomerRequest{
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
			},
			wantErr:    true,
			wantErrMsg: "db error",
		},
		{
			name: "Sucesso",
			req: request.NewCustomerRequest{
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
			},

			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			countryRepo := &mocks_repo.CountryRepoMock{}
			customerRepo := &mocks_repo.CustomerRepoMock{}
			dataManager := &mocks.DataManagerMock{
				CountryRepoField:  countryRepo,
				CustomerRepoField: customerRepo,
			}
			serviceManager := &mocks.ServiceManagerMock{
				DataManagerField: dataManager,
			}
			svc := service.NewCustomerService(serviceManager)
			resp, err := svc.Save(context.Background(), tc.req)

			if tc.wantErr {
				require.Error(t, err)
				if tc.wantErrMsg != "" {
					require.Contains(t, err.Error(), tc.wantErrMsg)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.req.Email, resp.Email)
				require.Equal(t, tc.req.Name, resp.Name)
			}
		})
	}
}
