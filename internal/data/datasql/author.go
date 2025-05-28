package datasql

import (
	"context"
	"fmt"
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type authorRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *authorRepository) Save(ctx context.Context, author entity.Author) (model.AuthorModel, error) {
	authorModel := model.ToAuthorModel(author)

	query := `
		INSERT INTO author (name, email, description, created_at)
		VALUES (?, ?, ?, ?)
	`

	now := time.Now()
	authorModel.CreatedAt = now

	result, err := r.conn.ExecContext(
		ctx,
		query,
		authorModel.Name,
		authorModel.Email,
		authorModel.Description,
		authorModel.CreatedAt,
	)
	if err != nil {
		return model.AuthorModel{}, fmt.Errorf("failed to insert author: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return model.AuthorModel{}, fmt.Errorf("failed to get inserted ID: %w", err)
	}
	authorModel.ID = uint(id)
	return authorModel, nil
}
