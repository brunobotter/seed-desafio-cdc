package contract

import (
	"context"

	"github.com/brunobotter/casa-codigo/configs/mapping"
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
}

type AuthorService interface {
	Save(ctx context.Context, request request.NewAuthorRequest) (response response.AuthorResponse, err error)
}
