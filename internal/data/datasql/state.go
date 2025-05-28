package datasql

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type stateRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *stateRepository) Save(ctx context.Context, country entity.State) (model.StateModel, error) {
	stateModel := model.ToStateModel(country)

	/*if err := r.db.WithContext(ctx).Create(&stateModel).Error; err != nil {
		return model.StateModel{}, err
	}*/

	return stateModel, nil
}
