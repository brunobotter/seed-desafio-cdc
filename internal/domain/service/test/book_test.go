package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
	"github.com/brunobotter/casa-codigo/internal/domain/service"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util/mocks"
	mocks_repo "github.com/brunobotter/casa-codigo/internal/util/mocks/repo"
	"github.com/stretchr/testify/require"
)

func TestBookService_Save(t *testing.T) {
	tests := []struct {
		name          string
		req           request.NewBookRequest
		setupRepoMock func() *mocks_repo.BookRepoMock
		wantErr       bool
		wantErrMsg    string
	}{
		{
			name:          "Titulo vazio",
			req:           request.NewBookRequest{Title: "", Resume: "fulano@email.com", Summary: "desc", Price: 100, Page: 100, ISBN: "12345678", PublishDate: time.Now()},
			setupRepoMock: func() *mocks_repo.BookRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "title is required",
		},
		{
			name:          "resumo vazio",
			req:           request.NewBookRequest{Title: "teste", Resume: "", Summary: "desc", Price: 100, Page: 100, ISBN: "12345678", PublishDate: time.Now()},
			setupRepoMock: func() *mocks_repo.BookRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "resume is required",
		},
		{
			name: "resumo maior que 500 caracteres",
			req: request.NewBookRequest{Title: "teste", Resume: "testetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetestetesteestetestetestetestetestetestetesteteste",
				Summary: "desc", Price: 100, Page: 100, ISBN: "12345678", PublishDate: time.Now()},
			setupRepoMock: func() *mocks_repo.BookRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "resume must be at most 500 characters",
		},
		{
			name:          "Preço menor que 20",
			req:           request.NewBookRequest{Title: "teste", Resume: "teste", Summary: "desc", Price: 10, Page: 100, ISBN: "12345678", PublishDate: time.Now()},
			setupRepoMock: func() *mocks_repo.BookRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "price must be at least 20",
		},
		{
			name:          "Pagina menor que 100",
			req:           request.NewBookRequest{Title: "teste", Resume: "teste", Summary: "desc", Price: 100, Page: 10, ISBN: "12345678", PublishDate: time.Now()},
			setupRepoMock: func() *mocks_repo.BookRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "page must be at least 100",
		},
		{
			name:          "ISBN vazio",
			req:           request.NewBookRequest{Title: "teste", Resume: "teste", Summary: "desc", Price: 100, Page: 100, ISBN: "", PublishDate: time.Now()},
			setupRepoMock: func() *mocks_repo.BookRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "isbn is required",
		},
		{
			name:          "data publicação e zero",
			req:           request.NewBookRequest{Title: "teste", Resume: "teste", Summary: "desc", Price: 100, Page: 100, ISBN: "123", PublishDate: time.Time{}},
			setupRepoMock: func() *mocks_repo.BookRepoMock { return nil },
			wantErr:       true,
			wantErrMsg:    "publish_date is required",
		},
		{
			name: "Sucesso",
			req:  request.NewBookRequest{Title: "teste", Resume: "teste", Summary: "desc", Price: 100, Page: 100, ISBN: "123", PublishDate: time.Now()},
			setupRepoMock: func() *mocks_repo.BookRepoMock {
				return &mocks_repo.BookRepoMock{
					SaveFunc: func(ctx context.Context, book entity.Book) (model.BookModel, error) {
						return model.BookModel{
							ID:          1,
							Title:       book.Title,
							Resume:      book.Resume,
							ISBN:        book.ISBN,
							Summary:     book.Summary,
							Price:       book.Price,
							Page:        book.Page,
							PublishDate: book.PublishDate,
						}, nil
					},
				}
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var repoMock contract.BookRepository
			if tc.setupRepoMock != nil {
				repoMock = tc.setupRepoMock()
			}
			mockDataManager := &mocks.DataManagerMock{
				BookRepoField: repoMock,
			}
			mockServiceManager := &mocks.ServiceManagerMock{
				DataManagerField: mockDataManager,
			}
			svc := service.NewBookService(mockServiceManager)
			resp, err := svc.Save(context.Background(), tc.req, 1, 1)

			if tc.wantErr {
				require.Error(t, err)
				if tc.wantErrMsg != "" {
					require.Contains(t, err.Error(), tc.wantErrMsg)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.req.Resume, resp.Resume)
				require.Equal(t, tc.req.Title, resp.Title)
				require.Equal(t, tc.req.ISBN, resp.ISBN)
				require.Equal(t, tc.req.Page, resp.Page)
				require.Equal(t, tc.req.Price, resp.Price)
				require.Equal(t, tc.req.Summary, resp.Summary)
				require.Equal(t, tc.req.PublishDate, resp.PublishDate)

			}
		})
	}
}
