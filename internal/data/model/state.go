package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type StateModel struct {
	ID        uint   `gorm:"primaryKey"`
	Statename string `gorm:"unique"`

	CreatedAt time.Time

	CountryID uint
	Country   CountryModel
}

func ToStateModel(e entity.State) StateModel {
	return StateModel{
		Statename: e.Name,
		CountryID: uint(e.CountryId),
	}
}
