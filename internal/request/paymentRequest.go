package request

import (
	"fmt"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type NewPaymentRequest struct {
	Total float64        `json:"total" validate:"required"`
	Itens []ItensRequest `json:"itens" validate:"required"`
}

type ItensRequest struct {
	BookId int64 `json:"book_id" validate:"required"`
	Amount int64 `json:"amount" validate:"required"`
}

func (r NewPaymentRequest) ToEntity() (entity.Payment, error) {
	if len(r.Itens) == 0 {
		return entity.Payment{}, fmt.Errorf("é necessário ao menos um item no pagamento")
	}
	var itens []entity.Itens
	for _, itemReq := range r.Itens {
		itens = append(itens, entity.Itens{
			BookId: itemReq.BookId,
			Amount: itemReq.Amount,
		})
	}

	return entity.Payment{
		Total: r.Total,
		Itens: itens,
	}, nil
}
