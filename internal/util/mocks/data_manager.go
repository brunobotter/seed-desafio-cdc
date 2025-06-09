package mocks

import "github.com/brunobotter/casa-codigo/internal/domain/contract"

type DataManagerMock struct {
	AuthorRepoField   contract.AuthorRepository
	CategoryRepoField contract.CategoryRepository
	BookRepoField     contract.BookRepository
}

func (m *DataManagerMock) AuthorRepo() contract.AuthorRepository     { return m.AuthorRepoField }
func (m *DataManagerMock) CategoryRepo() contract.CategoryRepository { return m.CategoryRepoField }
func (m *DataManagerMock) BookRepo() contract.BookRepository         { return m.BookRepoField }
func (m *DataManagerMock) CountryRepo() contract.CountryRepository   { return nil }
func (m *DataManagerMock) StateRepo() contract.StateRepository       { return nil }
func (m *DataManagerMock) CustomerRepo() contract.CustomerRepository { return nil }
func (m *DataManagerMock) PaymentRepo() contract.PaymentRepository   { return nil }
func (m *DataManagerMock) CouponRepo() contract.CouponRepository     { return nil }
