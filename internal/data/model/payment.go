package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type PaymentModel struct {
	ID         uint
	CustomerID int64
	Total      float64
	CreatedAt  time.Time
	Itens      []PaymentItemModel
}

type PaymentItemModel struct {
	BookID int64
	Amount int64
	Price  float64
}

func ToPaymentModel(p entity.Payment) PaymentModel {
	var itens []PaymentItemModel
	for _, item := range p.Itens {
		itens = append(itens, PaymentItemModel{
			BookID: item.BookId,
			Amount: item.Amount,
			Price:  item.Price,
		})
	}
	return PaymentModel{
		Total: p.Total,
		Itens: itens,
	}
}
