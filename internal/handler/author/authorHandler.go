package authorHandler

import (
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/gin-gonic/gin"
)

func SaveNewAuthor(ctx *gin.Context) {
	var request request.NewAuthorRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Dados inválidos: " + err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Autor criado com sucesso!",
	})
}
