package service_test

import (
	"context"
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

func TestCategoryService_Save(t *testing.T) {
	tests := []struct {
		name          string
		req           request.NewCategoryRequest
		setupRepoMock func() *mocks_repo.CategoryRepoMock
		wantErr       bool
		wantErrMsg    string
	}{
		{
			name:          "Categoria vazio",
			req:           request.NewCategoryRequest{Name: ""},
			setupRepoMock: func() *mocks_repo.CategoryRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "invalid category",
		},
		{
			name: "Sucesso",
			req:  request.NewCategoryRequest{Name: "ciencia"},
			setupRepoMock: func() *mocks_repo.CategoryRepoMock {
				return &mocks_repo.CategoryRepoMock{
					SaveFunc: func(ctx context.Context, category entity.Category) (model.CategoryModel, error) {
						return model.CategoryModel{
							ID:   1,
							Name: category.Name,
						}, nil
					},
				}
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var repoMock contract.CategoryRepository
			if tc.setupRepoMock != nil {
				repoMock = tc.setupRepoMock()
			}
			mockDataManager := &mocks.DataManagerMock{
				CategoryRepoField: repoMock,
			}
			mockServiceManager := &mocks.ServiceManagerMock{
				DataManagerField: mockDataManager,
			}
			svc := service.NewCategoryService(mockServiceManager)
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
