package contract

import (
	"context"

	"github.com/brunobotter/casa-codigo/configs/mapping"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type ServiceManager interface {
	Config() *mapping.Config
	DB() DataManager
	InternalService() InternalService
}

type InternalService interface {
	AuthorService() AuthorService
	CategoryService() CategoryService
	BookService() BookService
	CountryService() CountryService
	StateService() StateService
	PaymentService() PaymentService
	CustomerService() CustomerService
	CouponService() CouponService
}

type AuthorService interface {
	Save(ctx context.Context, request request.NewAuthorRequest) (response response.AuthorResponse, err error)
}

type CategoryService interface {
	Save(ctx context.Context, request request.NewCategoryRequest) (response response.CategoryResponse, err error)
}

type BookService interface {
	Save(ctx context.Context, request request.NewBookRequest, categoryId int64, authorId int64) (response response.BookResponse, err error)
	GetById(ctx context.Context, bookId int64) (response response.BookResponse, err error)
	GetAll(ctx context.Context) (response response.BookListResponse, err error)
}

type CountryService interface {
	Save(ctx context.Context, request request.NewCountryRequest) (response response.CountryResponse, err error)
}

type StateService interface {
	Save(ctx context.Context, request request.NewStateRequest, countryId int64) (response response.StateResponse, err error)
}

type PaymentService interface {
	Save(ctx context.Context, request request.NewPaymentRequest, customerId int64) (response response.PaymentResponse, err error)
	FillItemPrices(ctx context.Context, items []entity.Itens) error
	ApplyCoupon(ctx context.Context, payment *entity.Payment) error
}

type CustomerService interface {
	Save(ctx context.Context, request request.NewCustomerRequest) (response response.CustomerResponse, err error)
}

type CouponService interface {
	Save(ctx context.Context, request request.NewCouponRequest) (response response.CouponResponse, err error)
}
