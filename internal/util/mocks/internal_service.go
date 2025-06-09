package mocks

import "github.com/brunobotter/casa-codigo/internal/domain/contract"

type InternalServiceMock struct {
	AuthorServiceField   contract.AuthorService
	CategoryServiceField contract.CategoryService
	BookServiceField     contract.BookService
	CountryServiceField  contract.CountryService
	StateServiceField    contract.StateService
	PaymentServiceField  contract.PaymentService
	CustomerServiceField contract.CustomerService
	CouponServiceField   contract.CouponService
}

func NewInternalServiceMock() *InternalServiceMock {
	return &InternalServiceMock{}
}

func (m *InternalServiceMock) AuthorService() contract.AuthorService {
	return m.AuthorServiceField
}
func (m *InternalServiceMock) CategoryService() contract.CategoryService {
	return m.CategoryServiceField
}
func (m *InternalServiceMock) BookService() contract.BookService {
	return m.BookServiceField
}
func (m *InternalServiceMock) CountryService() contract.CountryService {
	return m.CountryServiceField
}
func (m *InternalServiceMock) StateService() contract.StateService {
	return m.StateServiceField
}
func (m *InternalServiceMock) PaymentService() contract.PaymentService {
	return m.PaymentServiceField
}
func (m *InternalServiceMock) CustomerService() contract.CustomerService {
	return m.CustomerServiceField
}
func (m *InternalServiceMock) CouponService() contract.CouponService {
	return m.CouponServiceField
}
