package handler

import (
	"github.com/brunobotter/casa-codigo/configs"
	"github.com/brunobotter/casa-codigo/internal/handler/controller"
)

var (
	AuthorController   *controller.AuthorController
	CategoryController *controller.CategoryController
)

func InitializeHandler(deps *configs.Deps) {
	AuthorController = controller.NewAuthorController(deps.Svc)
	CategoryController = controller.NewCategoryController(deps.Svc)
}
