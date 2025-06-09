package mocks

import (
	"context"

	"github.com/brunobotter/casa-codigo/configs"
	"github.com/brunobotter/casa-codigo/configs/mapping"
)

type Setup struct {
	InternalServiceMock *InternalServiceMock
	Deps                *configs.Deps
	Ctx                 context.Context
	Config              *mapping.Config
}

func NewSetup() *Setup {
	return &Setup{}
}

func (s *Setup) WithContext() *Setup {
	s.Ctx = context.Background()
	return s
}

func (s *Setup) WithConfig() *Setup {
	s.Config = &mapping.Config{}
	return s
}

func (s *Setup) WithInternalServices() *Setup {
	s.InternalServiceMock = NewInternalServiceMock()
	return s
}
