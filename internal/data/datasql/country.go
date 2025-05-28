package datasql

import (
	"context"

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

	/*if err := r.db.WithContext(ctx).Create(&countryModel).Error; err != nil {
		return model.CountryModel{}, err
	}*/

	return countryModel, nil
}
