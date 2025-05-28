package data

import (
	"github.com/brunobotter/casa-codigo/configs/mapping"
	"github.com/brunobotter/casa-codigo/internal/data/datasql"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
)

func Connect(cfg *mapping.Config) (contract.DataManager, error) {
	return datasql.Instance(cfg)
}
