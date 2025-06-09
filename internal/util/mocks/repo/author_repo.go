package mocks_repo

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type AuthorRepoMock struct {
	SaveFunc func(ctx context.Context, author entity.Author) (model.AuthorModel, error)
}

func (m *AuthorRepoMock) Save(ctx context.Context, author entity.Author) (model.AuthorModel, error) {
	return m.SaveFunc(ctx, author)
}
