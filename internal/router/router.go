package router

import (
	authorRouter "github.com/brunobotter/casa-codigo/internal/router/author"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	authorRouter.InitializeAuthorRouters(router)

	router.Run(":8080")
}
