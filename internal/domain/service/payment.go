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

func (s *paymentService) Save(ctx context.Context, request request.NewPaymentRequest, customerId int64) (response.PaymentResponse, error) {
	payment, err := request.ToEntity()
	if err != nil {
		return response.PaymentResponse{}, err
	}

	_, err = s.svc.DB().CustomerRepo().GetById(ctx, customerId)
	if err != nil {
		return response.PaymentResponse{}, err
	}

	payment.CustomerId = customerId

	err = s.svc.InternalService().PaymentService().FillItemPrices(ctx, payment.Itens)
	if err != nil {
		return response.PaymentResponse{}, err
	}

	err = verifyTotalPrice(payment)
	if err != nil {
		return response.PaymentResponse{}, err
	}

	paymentData, err := s.svc.DB().PaymentRepo().Save(ctx, payment)
	if err != nil {
		return response.PaymentResponse{}, err
	}

	return response.FromPaymentResponse(paymentData), nil
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

func (s *paymentService) FillItemPrices(ctx context.Context, items []entity.Itens) error {
	for i := range items {
		book, err := s.svc.DB().BookRepo().GetById(ctx, items[i].BookId)
		if err != nil {
			return err
		}
		items[i].Price = book.Price
	}
	return nil
}
