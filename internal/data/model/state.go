package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type StateModel struct {
	ID        uint
	Statename string
	CountryID uint
	CreatedAt time.Time
}

func ToStateModel(e entity.State) StateModel {
	return StateModel{
		Statename: e.Name,
		CountryID: uint(e.CountryId),
	}
}
