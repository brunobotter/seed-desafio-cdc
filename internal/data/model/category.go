package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type CategoryModel struct {
	ID        uint
	Name      string
	CreatedAt time.Time
}

func ToCategoryModel(e entity.Category) CategoryModel {
	return CategoryModel{
		Name: e.Name,
	}
}
