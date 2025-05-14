package datasql

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
	"gorm.io/gorm"
)

type authorRepository struct {
	data contract.RepoManager
	db   *gorm.DB
}

func (r *authorRepository) Save(ctx context.Context, author entity.Author) (model.AuthorModel, error) {
	authorModel := model.ToAuthorModel(author)

	if err := r.db.WithContext(ctx).Create(&authorModel).Error; err != nil {
		return model.AuthorModel{}, err
	}

	return authorModel, nil
}
