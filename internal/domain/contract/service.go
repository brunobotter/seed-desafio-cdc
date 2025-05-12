package contract

import (
	"context"

	"github.com/brunobotter/casa-codigo/configs/mapping"
	"github.com/brunobotter/casa-codigo/internal/request"
	"gorm.io/gorm"
)

type ServiceManager interface {
	Config() *mapping.Config
	DB() *gorm.DB
	InternalService() InternalService
}

type InternalService interface {
	AuthorService() AuthorService
}

type AuthorService interface {
	Save(ctx context.Context, request request.NewAuthorRequest) error
}
