package datasql

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type paymentRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *paymentRepository) Save(ctx context.Context, customer entity.Payment) (model.PaymentModel, error) {
	/*customerModel := model.ToCustomerModel(customer)

	query := `
		INSERT INTO customer (email, name, lastname, document, address, complement, city, country, state, phone, cep, created_at)
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
		return model.PaymentModel{}, fmt.Errorf("failed to insert customer: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.PaymentModel{}, fmt.Errorf("failed to get inserted ID: %w", err)
	}

	customerModel.ID = uint(id)

	return customerModel, nil*/
	return model.PaymentModel{}, nil
}
