package main

import (
	"github.com/brunobotter/casa-codigo/configs"
	"github.com/brunobotter/casa-codigo/internal/router"
)

var (
	logger *configs.Logger
)

func main() {
	logger = configs.GetLogger("main")
	deps := configs.Init().ConfigAll()
	if deps == nil || deps.DB == nil || deps.Svc == nil {
		logger.Errorf("Config initialize with error")
		return
	}
	router.Initialize(deps)
}
