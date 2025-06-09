package service_test

import (
	"context"
	"errors"
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

func TestBookService_GetById(t *testing.T) {
	tests := []struct {
		name          string
		bookId        int64
		setupRepoMock func() *mocks_repo.BookRepoMock
		wantErr       bool
		wantErrMsg    string
		wantModel     model.BookByIdModel // Para comparar no sucesso
	}{
		{
			name:   "ID inválido (boundary)",
			bookId: 0,
			setupRepoMock: func() *mocks_repo.BookRepoMock {
				return &mocks_repo.BookRepoMock{
					GetByIdFunc: func(ctx context.Context, id int64) (model.BookByIdModel, error) {
						return model.BookByIdModel{}, errors.New("invalid id")
					},
				}
			},
			wantErr:    true,
			wantErrMsg: "invalid id",
		},
		{
			name:   "Book não encontrado (erro do repo)",
			bookId: 99,
			setupRepoMock: func() *mocks_repo.BookRepoMock {
				return &mocks_repo.BookRepoMock{
					GetByIdFunc: func(ctx context.Context, id int64) (model.BookByIdModel, error) {
						return model.BookByIdModel{}, errors.New("not found")
					},
				}
			},
			wantErr:    true,
			wantErrMsg: "not found",
		},
		{
			name:   "Sucesso (book encontrado)",
			bookId: 1,
			setupRepoMock: func() *mocks_repo.BookRepoMock {
				return &mocks_repo.BookRepoMock{
					GetByIdFunc: func(ctx context.Context, id int64) (model.BookByIdModel, error) {
						return model.BookByIdModel{
							ID:          1,
							Title:       "Livro Teste",
							Resume:      "teste",
							Summary:     "desc",
							Price:       100,
							Page:        100,
							ISBN:        "123",
							PublishDate: time.Now(),
						}, nil
					},
				}
			},
			wantErr: false,
			wantModel: model.BookByIdModel{
				ID:          1,
				Title:       "Livro Teste",
				Resume:      "teste",
				Summary:     "desc",
				Price:       100,
				Page:        100,
				ISBN:        "123",
				PublishDate: time.Now(),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := tc.setupRepoMock()
			mockDataManager := &mocks.DataManagerMock{BookRepoField: repoMock}
			mockServiceManager := &mocks.ServiceManagerMock{DataManagerField: mockDataManager}
			svc := service.NewBookService(mockServiceManager)

			resp, err := svc.GetById(context.Background(), tc.bookId)
			if tc.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.wantErrMsg)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.wantModel.ID, resp.ID)
				require.Equal(t, tc.wantModel.Title, resp.Title)
				require.Equal(t, tc.wantModel.Resume, resp.Resume)
				require.Equal(t, tc.wantModel.Summary, resp.Summary)
				require.Equal(t, tc.wantModel.Price, resp.Price)
				require.Equal(t, tc.wantModel.Page, resp.Page)
				require.Equal(t, tc.wantModel.ISBN, resp.ISBN)
				require.Equal(t, tc.wantModel.PublishDate, resp.PublishDate)
			}
		})
	}
}

func TestBookService_GetAll(t *testing.T) {

	tests := []struct {
		name          string
		setupRepoMock func() *mocks_repo.BookRepoMock
		wantErr       bool
		wantErrMsg    string
		wantBooks     []model.BookModel
	}{
		{
			name: "Sem livros (boundary)",
			setupRepoMock: func() *mocks_repo.BookRepoMock {
				return &mocks_repo.BookRepoMock{
					GetAllFunc: func(ctx context.Context) ([]model.BookByAllModel, error) {
						return []model.BookByAllModel{}, nil
					},
				}
			},
			wantErr:   false,
			wantBooks: []model.BookModel{},
		},
		{
			name: "Erro no repo",
			setupRepoMock: func() *mocks_repo.BookRepoMock {
				return &mocks_repo.BookRepoMock{
					GetAllFunc: func(ctx context.Context) ([]model.BookByAllModel, error) {
						return nil, errors.New("repo error")
					},
				}
			},
			wantErr:    true,
			wantErrMsg: "repo error",
		},
		{
			name: "Sucesso (um livro)",
			setupRepoMock: func() *mocks_repo.BookRepoMock {
				return &mocks_repo.BookRepoMock{
					GetAllFunc: func(ctx context.Context) ([]model.BookByAllModel, error) {
						return []model.BookByAllModel{
							{
								ID:    1,
								Title: "Livro Um",
							},
						}, nil
					},
				}
			},
			wantErr: false,
			wantBooks: []model.BookModel{
				{
					ID:    1,
					Title: "Livro Um",
				},
			},
		},
		{
			name: "Sucesso (vários livros)",
			setupRepoMock: func() *mocks_repo.BookRepoMock {
				return &mocks_repo.BookRepoMock{
					GetAllFunc: func(ctx context.Context) ([]model.BookByAllModel, error) {
						return []model.BookByAllModel{
							{
								ID:    1,
								Title: "Livro 1",
							},
							{
								ID:    2,
								Title: "Livro 2",
							},
						}, nil
					},
				}
			},
			wantErr: false,
			wantBooks: []model.BookModel{
				{
					ID:    1,
					Title: "Livro 1",
				},
				{
					ID:    2,
					Title: "Livro 2",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := tc.setupRepoMock()
			dataManager := &mocks.DataManagerMock{BookRepoField: repoMock}
			svcManager := &mocks.ServiceManagerMock{DataManagerField: dataManager}
			svc := service.NewBookService(svcManager)

			books, err := svc.GetAll(context.Background())

			if tc.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.wantErrMsg)
			} else {
				require.NoError(t, err)
				require.Equal(t, len(tc.wantBooks), len(books))
				for i, wantBook := range tc.wantBooks {
					gotBook := books[i]
					require.Equal(t, wantBook.ID, gotBook.ID)
					require.Equal(t, wantBook.Title, gotBook.Title)
				}
			}
		})
	}
}
