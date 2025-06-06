package datasql

import (
	"context"
	"database/sql"
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

func (r *customerRepository) GetById(ctx context.Context, customerId int64) (model.CustomerModel, error) {
	query := `
		SELECT 
			id, email, first_name, last_name, document, address, complement, city,
			country_id, state_id, phone, cep, created_at
		FROM customer
		WHERE id = ?
	`

	var customer model.CustomerModel
	var stateID sql.NullInt64

	err := r.conn.QueryRowContext(ctx, query, customerId).Scan(
		&customer.ID,
		&customer.Email,
		&customer.Name,
		&customer.Lastname,
		&customer.Document,
		&customer.Address,
		&customer.Complement,
		&customer.City,
		&customer.Country,
		&stateID,
		&customer.Phone,
		&customer.CEP,
		&customer.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.CustomerModel{}, fmt.Errorf("customer with id %d not found", customerId)
		}
		return model.CustomerModel{}, fmt.Errorf("failed to get customer by id: %w", err)
	}

	if stateID.Valid {
		state := int64(stateID.Int64)
		customer.State = &state
	} else {
		customer.State = nil
	}

	return customer, nil
}
