package datasql

import (
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"gorm.io/gorm"
)

type dataManager struct {
	db *gorm.DB
}

func NewDataManager(db *gorm.DB) contract.DataManager {
	return &dataManager{db: db}
}

func (d *dataManager) DB() *gorm.DB {
	return d.db
}

func (d *dataManager) AuthorRepo() contract.AuthorRepository {
	return &authorRepository{db: d.db}
}
