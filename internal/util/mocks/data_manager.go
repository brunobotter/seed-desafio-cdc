package mocks

import "github.com/brunobotter/casa-codigo/internal/domain/contract"

type DataManagerMock struct {
	AuthorRepoField   contract.AuthorRepository
	CategoryRepoField contract.CategoryRepository
	BookRepoField     contract.BookRepository
	CountryRepoField  contract.CountryRepository
	StateRepoField    contract.StateRepository
	CustomerRepoField contract.CustomerRepository
	PaymentRepoField  contract.PaymentRepository
	CouponRepoField   contract.CouponRepository
}

func (m *DataManagerMock) AuthorRepo() contract.AuthorRepository     { return m.AuthorRepoField }
func (m *DataManagerMock) CategoryRepo() contract.CategoryRepository { return m.CategoryRepoField }
func (m *DataManagerMock) BookRepo() contract.BookRepository         { return m.BookRepoField }
func (m *DataManagerMock) CountryRepo() contract.CountryRepository   { return m.CountryRepoField }
func (m *DataManagerMock) StateRepo() contract.StateRepository       { return m.StateRepoField }
func (m *DataManagerMock) CustomerRepo() contract.CustomerRepository { return m.CustomerRepoField }
func (m *DataManagerMock) PaymentRepo() contract.PaymentRepository   { return m.PaymentRepoField }
func (m *DataManagerMock) CouponRepo() contract.CouponRepository     { return m.CouponRepoField }
