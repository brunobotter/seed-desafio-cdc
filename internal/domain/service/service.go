package service

import (
	"github.com/brunobotter/casa-codigo/configs/mapping"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
)

type serviceManager struct {
	cfg             *mapping.Config
	db              contract.DataManager
	internalService contract.InternalService
}

func (s *serviceManager) Config() *mapping.Config {
	return s.cfg
}

func (s *serviceManager) DB() contract.DataManager {
	return s.db
}
func (s *serviceManager) InternalService() contract.InternalService {
	return s.internalService
}

type internalServices struct {
	authorService   contract.AuthorService
	categoryService contract.CategoryService
	bookService     contract.BookService
	countryService  contract.CountryService
	stateService    contract.StateService
	paymentService  contract.PaymentService
	customerService contract.CustomerService
	couponService   contract.CouponService
}

func (s *internalServices) AuthorService() contract.AuthorService {
	return s.authorService
}

func (s *internalServices) CategoryService() contract.CategoryService {
	return s.categoryService
}

func (s *internalServices) BookService() contract.BookService {
	return s.bookService
}

func (s *internalServices) CountryService() contract.CountryService {
	return s.countryService
}

func (s *internalServices) StateService() contract.StateService {
	return s.stateService
}

func (s *internalServices) PaymentService() contract.PaymentService {
	return s.paymentService
}

func (s *internalServices) CustomerService() contract.CustomerService {
	return s.customerService
}

func (s *internalServices) CouponService() contract.CouponService {
	return s.couponService
}

type ServiceDeps struct {
	Cfg *mapping.Config
	DB  contract.DataManager
}

func New(deps ServiceDeps) (contract.ServiceManager, error) {
	instance := &serviceManager{
		db:  deps.DB,
		cfg: deps.Cfg,
	}

	internalServices := internalServices{}
	internalServices.authorService = NewAuthorService(instance)
	internalServices.categoryService = NewCategoryService(instance)
	internalServices.bookService = NewBookService(instance)
	internalServices.countryService = NewCountryService(instance)
	internalServices.stateService = NewStateService(instance)
	internalServices.paymentService = NewPaymentService(instance)
	internalServices.customerService = NewCustomerService(instance)
	internalServices.couponService = NewCouponService(instance)

	instance.internalService = &internalServices
	return instance, nil
}
