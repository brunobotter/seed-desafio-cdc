package handler

import (
	"github.com/brunobotter/casa-codigo/configs"
	"github.com/brunobotter/casa-codigo/internal/handler/controller"
)

var (
	AuthorController   *controller.AuthorController
	CategoryController *controller.CategoryController
	BookController     *controller.BookController
	CountryController  *controller.CountryController
	StateController    *controller.StateController
	PaymentController  *controller.PaymentController
	CustomerController *controller.CustomerController
	CouponController   *controller.CouponController
)

func InitializeHandler(deps *configs.Deps) {
	AuthorController = controller.NewAuthorController(deps.Svc)
	CategoryController = controller.NewCategoryController(deps.Svc)
	BookController = controller.NewBookController(deps.Svc)
	CountryController = controller.NewCountryController(deps.Svc)
	StateController = controller.NewStateController(deps.Svc)
	PaymentController = controller.NewPaymentController(deps.Svc)
	CustomerController = controller.NewCustomerController(deps.Svc)
	CouponController = controller.NewCouponController(deps.Svc)
}
