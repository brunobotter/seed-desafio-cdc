package router

import (
	"github.com/brunobotter/casa-codigo/configs"
	"github.com/brunobotter/casa-codigo/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitializeCouponRouters(router *gin.Engine, deps *configs.Deps) {
	handler.InitializeHandler(deps)
	v1 := router.Group("api/v1/coupons")
	{
		v1.POST("/save", handler.CouponController.SaveNewCoupon)
	}
}
