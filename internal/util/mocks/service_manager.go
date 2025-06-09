package mocks

import (
	"github.com/brunobotter/casa-codigo/configs/mapping"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
)

type ServiceManagerMock struct {
	ConfigField          *mapping.Config
	DataManagerField     contract.DataManager
	InternalServiceField contract.InternalService
}

func (m *ServiceManagerMock) Config() *mapping.Config {
	return m.ConfigField
}

func (m *ServiceManagerMock) DB() contract.DataManager {
	return m.DataManagerField
}

func (m *ServiceManagerMock) InternalService() contract.InternalService {
	return m.InternalServiceField
}
