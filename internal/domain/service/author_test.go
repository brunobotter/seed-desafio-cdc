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

func TestAuthorService_Save(t *testing.T) {
	tests := []struct {
		name          string
		req           request.NewAuthorRequest
		setupRepoMock func() *mocks_repo.AuthorRepoMock
		wantErr       bool
		wantErrMsg    string
	}{
		{
			name:          "Nome vazio",
			req:           request.NewAuthorRequest{Name: "", Email: "fulano@email.com", Description: "desc"},
			setupRepoMock: func() *mocks_repo.AuthorRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "invalid name",
		},
		{
			name:          "Descrição vazio",
			req:           request.NewAuthorRequest{Name: "fulano", Email: "fulano@email.com", Description: ""},
			setupRepoMock: func() *mocks_repo.AuthorRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "invalid description",
		},
		{
			name:          "Descrição maior que 400 caracteres",
			req:           request.NewAuthorRequest{Name: "fulano", Email: "fulano@email.com", Description: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			setupRepoMock: func() *mocks_repo.AuthorRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "invalid description",
		},
		{
			name:          "Email inválido (boundary/MC-DC)",
			req:           request.NewAuthorRequest{Name: "Fulano", Email: "invalid", Description: "desc"},
			setupRepoMock: func() *mocks_repo.AuthorRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "invalid email",
		},
		{
			name: "Repo retorna erro (decisão do repo)",
			req:  request.NewAuthorRequest{Name: "Fulano", Email: "fulano@email.com", Description: "desc"},
			setupRepoMock: func() *mocks_repo.AuthorRepoMock {
				return &mocks_repo.AuthorRepoMock{
					SaveFunc: func(ctx context.Context, author entity.Author) (model.AuthorModel, error) {
						return model.AuthorModel{}, errors.New("db error")
					},
				}
			},
			wantErr:    true,
			wantErrMsg: "db error",
		},
		{
			name: "Sucesso",
			req:  request.NewAuthorRequest{Name: "Fulano", Email: "fulano@email.com", Description: "desc"},
			setupRepoMock: func() *mocks_repo.AuthorRepoMock {
				return &mocks_repo.AuthorRepoMock{
					SaveFunc: func(ctx context.Context, author entity.Author) (model.AuthorModel, error) {
						return model.AuthorModel{
							ID:          1,
							Name:        author.Name,
							Email:       author.Email,
							Description: author.Description,
						}, nil
					},
				}
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var repoMock contract.AuthorRepository
			if tc.setupRepoMock != nil {
				repoMock = tc.setupRepoMock()
			}
			mockDataManager := &mocks.DataManagerMock{
				AuthorRepoField: repoMock,
			}
			mockServiceManager := &mocks.ServiceManagerMock{
				DataManagerField: mockDataManager,
			}
			svc := service.NewAuthorService(mockServiceManager)
			resp, err := svc.Save(context.Background(), tc.req)

			if tc.wantErr {
				require.Error(t, err)
				if tc.wantErrMsg != "" {
					require.Contains(t, err.Error(), tc.wantErrMsg)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.req.Name, resp.Name)
				require.Equal(t, tc.req.Email, resp.Email)
			}
		})
	}
}
