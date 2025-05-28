package datasql

import (
	"context"
	"fmt"
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type stateRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *stateRepository) Save(ctx context.Context, state entity.State) (model.StateModel, error) {
	stateModel := model.ToStateModel(state)

	query := `
		INSERT INTO state (name, country_id, created_at)
		VALUES (?, ?, ?)
	`

	now := time.Now()
	stateModel.CreatedAt = now

	result, err := r.conn.ExecContext(
		ctx,
		query,
		stateModel.Statename,
		stateModel.CountryId,
		stateModel.CreatedAt,
	)
	if err != nil {
		return model.StateModel{}, fmt.Errorf("failed to insert state: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return model.StateModel{}, fmt.Errorf("failed to get inserted ID: %w", err)
	}
	stateModel.ID = uint(id)
	return stateModel, nil
}
