package datasql

import (
	"context"
	"fmt"
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type customerRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *customerRepository) Save(ctx context.Context, customer entity.Customer) (model.CustomerModel, error) {
	customerModel := model.ToCustomerModel(customer)

	query := `
		INSERT INTO customer (email, first_name, last_name, document, address, complement, city, country_id, state_id, phone, cep, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	customerModel.CreatedAt = now

	result, err := r.conn.ExecContext(
		ctx,
		query,
		customerModel.Email,
		customerModel.Name,
		customerModel.Lastname,
		customerModel.Document,
		customerModel.Address,
		customerModel.Complement,
		customerModel.City,
		customerModel.Country,
		customerModel.State,
		customerModel.Phone,
		customerModel.CEP,
		customerModel.CreatedAt,
	)
	if err != nil {
		return model.CustomerModel{}, fmt.Errorf("failed to insert customer: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.CustomerModel{}, fmt.Errorf("failed to get inserted ID: %w", err)
	}

	customerModel.ID = uint(id)

	return customerModel, nil
}
