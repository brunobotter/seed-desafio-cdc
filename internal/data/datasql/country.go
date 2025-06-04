package datasql

import (
	"context"
	"database/sql"
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

func (r *countryRepository) VerifyCountryState(ctx context.Context, country string, state *string) (model.CountryState, error) {
	var query string
	var args []interface{}

	if state != nil {
		query = `
			SELECT s.id, c.id
			FROM state s
			JOIN country c ON s.country_id = c.id
			WHERE c.name = ? AND s.name = ?;
		`
		args = append(args, country, *state)
	} else {
		query = `
			SELECT s.id, c.id
			FROM state s
			JOIN country c ON s.country_id = c.id
			WHERE c.name = ?;
		`
		args = append(args, country)
	}

	row := r.conn.QueryRowContext(ctx, query, args...)
	fmt.Printf("test %v", row)
	var cs model.CountryState
	err := row.Scan(&cs.Stateid, &cs.Countryid)
	if err != nil {
		if err == sql.ErrNoRows {
			if state != nil {
				return model.CountryState{}, fmt.Errorf("state '%s' not found in country '%s'", *state, country)
			}
			return model.CountryState{}, fmt.Errorf("no state found for this country'%s'", country)
		}
		return model.CountryState{}, fmt.Errorf("error to verify state: %w", err)
	}

	return cs, nil
}
