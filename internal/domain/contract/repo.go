package contract

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	entity "github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type RepoManager interface {
	AuthorRepo() AuthorRepository
}

type AuthorRepository interface {
	Save(ctx context.Context, author entity.Author) (model.AuthorModel, error)
}
