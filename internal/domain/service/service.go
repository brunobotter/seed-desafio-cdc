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
	authorService contract.AuthorService
}

func (s *internalServices) AuthorService() contract.AuthorService {
	return s.authorService
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
	internalServices.authorService = NewAuthorInternalService(instance)
	instance.internalService = &internalServices
	return instance, nil
}
