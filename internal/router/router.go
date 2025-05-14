package router

import (
	"github.com/brunobotter/casa-codigo/configs"
	authorRouter "github.com/brunobotter/casa-codigo/internal/router/author"
	"github.com/gin-gonic/gin"
)

func Initialize(deps *configs.Deps) {
	router := gin.Default()

	authorRouter.InitializeAuthorRouters(router, deps)

	router.Run(":8080")
}
