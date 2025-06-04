package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type CountryModel struct {
	ID        uint
	Name      string
	CreatedAt time.Time
}

func ToCountryModel(e entity.Country) CountryModel {
	return CountryModel{
		Name: e.Name,
	}
}

type CountryState struct {
	Countryid int64
	Stateid   *int64
}
