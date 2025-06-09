package mocks

import "github.com/brunobotter/casa-codigo/internal/domain/contract"

type DataManagerMock struct {
	AuthorRepoField contract.AuthorRepository
}

func (m *DataManagerMock) AuthorRepo() contract.AuthorRepository {
	return m.AuthorRepoField
}

// Implemente os outros métodos, retornando nil ou mocks, se necessário
func (m *DataManagerMock) CategoryRepo() contract.CategoryRepository { return nil }
func (m *DataManagerMock) BookRepo() contract.BookRepository         { return nil }
func (m *DataManagerMock) CountryRepo() contract.CountryRepository   { return nil }
func (m *DataManagerMock) StateRepo() contract.StateRepository       { return nil }
func (m *DataManagerMock) CustomerRepo() contract.CustomerRepository { return nil }
func (m *DataManagerMock) PaymentRepo() contract.PaymentRepository   { return nil }
func (m *DataManagerMock) CouponRepo() contract.CouponRepository     { return nil }
