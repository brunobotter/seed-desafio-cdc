package handler

import (
	"github.com/brunobotter/casa-codigo/configs"
	"gorm.io/gorm"
)

var (
	logger *configs.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = configs.GetLogger("handler")
	db = configs.GetMySql()
	//service.InitializeService(db)
}
