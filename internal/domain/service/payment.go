package service

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
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

func (s *paymentService) Save(ctx context.Context, request request.NewPaymentRequest) (customerResponse response.PaymentResponse, err error) {

	return customerResponse, nil
}
