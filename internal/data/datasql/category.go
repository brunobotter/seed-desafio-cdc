package datasql

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type categoryRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *categoryRepository) Save(ctx context.Context, category entity.Category) (model.CategoryModel, error) {
	categoryModel := model.ToCategoryModel(category)

	/*if err := r.db.WithContext(ctx).Create(&categoryModel).Error; err != nil {
		return model.CategoryModel{}, err
	}*/

	return categoryModel, nil
}
