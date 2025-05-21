package router

import (
	"github.com/brunobotter/casa-codigo/configs"
	"github.com/brunobotter/casa-codigo/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitializeBookRouter(router *gin.Engine, deps *configs.Deps) {
	handler.InitializeHandler(deps)
	v1 := router.Group("api/v1/book")
	{
		v1.POST("save/:categoryId/:authorId", handler.BookController.SaveNewBook)
		v1.GET("/:bookId", handler.BookController.GetById)
	}
}
