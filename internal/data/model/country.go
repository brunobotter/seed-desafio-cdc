package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type CountryModel struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
}

func ToCountryModel(e entity.Country) CountryModel {
	return CountryModel{
		Name: e.Name,
	}
}
