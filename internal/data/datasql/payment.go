package datasql

import (
	"context"
	"fmt"
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type paymentRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *paymentRepository) Save(ctx context.Context, payment entity.Payment) (model.PaymentModel, error) {
	paymentModel := model.ToPaymentModel(payment)

	paymentModel.CreatedAt = time.Now()

	// Inserção na tabela payments
	insertPaymentQuery := `
		INSERT INTO payments (customer_id, total, created_at)
		VALUES (?, ?, ?)
	`
	result, err := r.conn.ExecContext(
		ctx,
		insertPaymentQuery,
		paymentModel.CustomerID,
		paymentModel.Total,
		paymentModel.CreatedAt,
	)
	if err != nil {
		return model.PaymentModel{}, fmt.Errorf("failed to insert payment: %w", err)
	}

	paymentID, err := result.LastInsertId()
	if err != nil {
		return model.PaymentModel{}, fmt.Errorf("failed to get inserted payment ID: %w", err)
	}
	paymentModel.ID = uint(paymentID)

	// Inserção dos itens
	insertItemQuery := `
		INSERT INTO payment_items (payment_id, book_id, amount, price)
		VALUES (?, ?, ?, ?)
	`

	for _, item := range paymentModel.Itens {
		_, err := r.conn.ExecContext(
			ctx,
			insertItemQuery,
			paymentID,
			item.BookID,
			item.Amount,
			item.Price,
		)
		if err != nil {
			return model.PaymentModel{}, fmt.Errorf("failed to insert payment item: %w", err)
		}
	}

	return paymentModel, nil
}
