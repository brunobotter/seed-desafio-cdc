package handler

import (
	"github.com/brunobotter/casa-codigo/configs"
	authorHandler "github.com/brunobotter/casa-codigo/internal/handler/author"
)

var (
	logger           *configs.Logger
	AuthorController *authorHandler.AuthorController
)

func InitializeHandler(deps *configs.Deps) {
	logger = configs.GetLogger("handler")
	AuthorController = authorHandler.NewAuthorController(deps.Cfg, deps.Svc)
}
