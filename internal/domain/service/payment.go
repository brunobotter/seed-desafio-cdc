package service

import (
	"context"
	"errors"
	"math"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type paymentService struct {
	svc contract.ServiceManager
}

func NewPaymentService(svc contract.ServiceManager) contract.PaymentService {
	return &paymentService{
		svc: svc,
	}
}

func (s *paymentService) Save(ctx context.Context, request request.NewPaymentRequest, customerId int64) (paymentResponse response.PaymentResponse, err error) {
	payment, err := request.ToEntity()
	_, err = s.svc.DB().CustomerRepo().GetById(ctx, customerId)
	payment.CustomerId = customerId

	if err != nil {
		return paymentResponse, err
	}
	for i := range payment.Itens {
		book, err := s.svc.DB().BookRepo().GetById(ctx, payment.Itens[i].BookId)
		if err != nil {
			return paymentResponse, err
		}

		payment.Itens[i].Price = book.Price
	}
	err = verifyTotalPrice(payment)
	if err != nil {
		return paymentResponse, err
	}

	paymentData, err := s.svc.DB().PaymentRepo().Save(ctx, payment)
	if err != nil {
		return paymentResponse, err
	}

	response := response.FromPaymentResponse(paymentData)
	return response, nil
}

func verifyTotalPrice(payment entity.Payment) error {
	var total float64

	for _, item := range payment.Itens {
		total += float64(item.Amount) * item.Price
	}

	const epsilon = 0.01

	if math.Abs(total-payment.Total) > epsilon {
		return errors.New("total does not match calculated total")
	}

	return nil
}
