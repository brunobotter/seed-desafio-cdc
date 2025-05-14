package contract

import (
	"gorm.io/gorm"
)

type DataManager interface {
	RepoManager
	DB() *gorm.DB
}
