package datasql

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
	"gorm.io/gorm"
)

type bookRepository struct {
	data contract.RepoManager
	db   *gorm.DB
}

func (r *bookRepository) Save(ctx context.Context, book entity.Book) (model.BookModel, error) {
	bookModel := model.ToBookModel(book)

	if err := r.db.WithContext(ctx).Create(&bookModel).Error; err != nil {
		return model.BookModel{}, err
	}

	return bookModel, nil
}
