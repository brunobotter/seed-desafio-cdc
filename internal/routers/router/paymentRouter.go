package router

import (
	"github.com/brunobotter/casa-codigo/configs"
	"github.com/brunobotter/casa-codigo/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitializePaymentRouters(router *gin.Engine, deps *configs.Deps) {
	handler.InitializeHandler(deps)
	v1 := router.Group("api/v1/payment")
	{
		v1.POST("/save/:customerId", handler.PaymentController.SaveNewPayment)
	}
}
