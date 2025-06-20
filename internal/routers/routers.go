package routers

import (
	"github.com/brunobotter/casa-codigo/configs"
	"github.com/brunobotter/casa-codigo/internal/routers/router"
	"github.com/gin-gonic/gin"
)

func Initialize(deps *configs.Deps) {
	gin := gin.Default()

	router.InitializeAuthorRouters(gin, deps)
	router.InitializeCategoryRouters(gin, deps)
	router.InitializeBookRouter(gin, deps)
	router.InitializeCountryRouters(gin, deps)
	router.InitializeStateRouters(gin, deps)
	router.InitializeCustomerRouters(gin, deps)
	router.InitializePaymentRouters(gin, deps)
	router.InitializeCouponRouters(gin, deps)

	gin.Run(":8080")
}
