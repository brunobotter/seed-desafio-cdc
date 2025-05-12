package authorRouter

import (
	"github.com/brunobotter/casa-codigo/internal/handler"
	authorHandler "github.com/brunobotter/casa-codigo/internal/handler/author"
	"github.com/gin-gonic/gin"
)

func InitializeAuthorRouters(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("api/v1/author")
	{
		v1.POST("/save", authorHandler.SaveNewAuthor)
	}
}
