package datasql

import (
	"context"
	"fmt"
	"time"

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

	query := `
		INSERT INTO category (name, created_at)
		VALUES (?, ?)
	`

	now := time.Now()
	categoryModel.CreatedAt = now

	result, err := r.conn.ExecContext(
		ctx,
		query,
		categoryModel.Name,
		categoryModel.CreatedAt,
	)
	if err != nil {
		return model.CategoryModel{}, fmt.Errorf("failed to insert category: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return model.CategoryModel{}, fmt.Errorf("failed to get inserted ID: %w", err)
	}
	categoryModel.ID = uint(id)
	return categoryModel, nil
}
