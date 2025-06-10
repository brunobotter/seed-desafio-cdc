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

func TestCountryService_Save(t *testing.T) {
	tests := []struct {
		name          string
		req           request.NewCountryRequest
		setupRepoMock func() *mocks_repo.CountryRepoMock
		wantErr       bool
		wantErrMsg    string
	}{
		{
			name:          "Nome vazio",
			req:           request.NewCountryRequest{Name: ""},
			setupRepoMock: func() *mocks_repo.CountryRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "invalid country",
		},
		{
			name: "Repo retorna erro (decisão do repo)",
			req:  request.NewCountryRequest{Name: "test"},
			setupRepoMock: func() *mocks_repo.CountryRepoMock {
				return &mocks_repo.CountryRepoMock{
					SaveFunc: func(ctx context.Context, country entity.Country) (model.CountryModel, error) {
						return model.CountryModel{}, errors.New("db error")
					},
				}
			},
			wantErr:    true,
			wantErrMsg: "db error",
		},
		{
			name: "Sucesso",
			req:  request.NewCountryRequest{Name: "Fulano"},
			setupRepoMock: func() *mocks_repo.CountryRepoMock {
				return &mocks_repo.CountryRepoMock{
					SaveFunc: func(ctx context.Context, country entity.Country) (model.CountryModel, error) {
						return model.CountryModel{
							ID:   1,
							Name: country.Name,
						}, nil
					},
				}
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var repoMock contract.CountryRepository
			if tc.setupRepoMock != nil {
				repoMock = tc.setupRepoMock()
			}
			mockDataManager := &mocks.DataManagerMock{
				CountryRepoField: repoMock,
			}
			mockServiceManager := &mocks.ServiceManagerMock{
				DataManagerField: mockDataManager,
			}
			svc := service.NewCountryService(mockServiceManager)
			resp, err := svc.Save(context.Background(), tc.req)

			if tc.wantErr {
				require.Error(t, err)
				if tc.wantErrMsg != "" {
					require.Contains(t, err.Error(), tc.wantErrMsg)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.req.Name, resp.Name)
			}
		})
	}
}
