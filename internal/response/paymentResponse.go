package response

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/util"
)

type PaymentResponse struct {
	ID         uint                  `json:"id"`
	CustomerID int64                 `json:"customer_id"`
	Total      float64               `json:"total"`
	Discount   float64               `json:"discount"`
	CreatedAt  time.Time             `json:"created_at"`
	Itens      []PaymentItemResponse `json:"itens"`
}

type PaymentItemResponse struct {
	BookID int64   `json:"book_id"`
	Amount int64   `json:"amount"`
	Price  float64 `json:"price"`
}

func FromPaymentResponse(m model.PaymentModel, discount float64) PaymentResponse {
	var itens []PaymentItemResponse
	for _, item := range m.Itens {
		itens = append(itens, PaymentItemResponse{
			BookID: item.BookID,
			Amount: item.Amount,
			Price:  item.Price,
		})
	}

	return PaymentResponse{
		ID:         m.ID,
		CustomerID: m.CustomerID,
		Total:      util.RoundToTwoDecimals(m.Total),
		Discount:   util.RoundToTwoDecimals(discount),
		CreatedAt:  m.CreatedAt,
		Itens:      itens,
	}
}
