package datasql

import (
	"context"
	"fmt"
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type countryRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *countryRepository) Save(ctx context.Context, country entity.Country) (model.CountryModel, error) {
	countryModel := model.ToCountryModel(country)

	query := `
		INSERT INTO country (name, created_at)
		VALUES (?, ?)
	`

	now := time.Now()
	countryModel.CreatedAt = now

	result, err := r.conn.ExecContext(
		ctx,
		query,
		countryModel.Name,
		countryModel.CreatedAt,
	)
	if err != nil {
		return model.CountryModel{}, fmt.Errorf("failed to insert country: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return model.CountryModel{}, fmt.Errorf("failed to get inserted ID: %w", err)
	}
	countryModel.ID = uint(id)
	return countryModel, nil
}
