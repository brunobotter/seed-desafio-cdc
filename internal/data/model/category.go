package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type CategoryModel struct {
	CategoryID uint   `gorm:"primaryKey"`
	Name       string `gorm:"unique"`
	CreatedAt  time.Time
}

func ToCategoryModel(e entity.Category) CategoryModel {
	return CategoryModel{
		Name: e.Name,
	}
}
