package service

import (
	"context"
	"errors"

	"payment-momo/internal/entity"
	"payment-momo/internal/model"
	"payment-momo/internal/persistence"
)

type PaymentService struct {
	momo *persistence.MomoClient
	repo persistence.OrderRepo
}

func NewPaymentService(m *persistence.MomoClient, r persistence.OrderRepo) *PaymentService {
	return &PaymentService{momo: m, repo: r}
}

func (s *PaymentService) CreatePayment(ctx context.Context, req model.CreatePaymentRequest) (*model.CreatePaymentResponse, error) {
	if req.Amount <= 0 {
		return nil, errors.New("amount must be > 0")
	}
	order := entity.NewOrder(req.Amount, req.Description)
	if err := s.repo.Save(order); err != nil {
		return nil, err
	}
	requestID := persistence.NewRequestID()
	payload := persistence.NewMomoCreatePayload(s.momo.Config(), order.ID, order.Amount, order.Description, requestID)
	momoRes, err := s.momo.CreatePayment(payload)
	if err != nil {
		return nil, err
	}
	return &model.CreatePaymentResponse{
		OrderID:   momoRes.OrderId,
		RequestID: momoRes.RequestId,
		PayURL:    momoRes.PayUrl,
		Message:   momoRes.Message,
		Code:      momoRes.ResultCode,
	}, nil
}

func (s *PaymentService) HandleIPN(ctx context.Context, cb model.MomoCallback) error {
	if !s.momo.VerifyCallbackSignature(cb) {
		return errors.New("invalid signature")
	}
	return s.repo.UpdateStatus(cb.OrderId, cb.ResultCode)
}

func (s *PaymentService) CheckStatus(ctx context.Context, orderID, requestID string) (*model.QueryStatusResponse, error) {
	return s.momo.QueryStatus(orderID, requestID)
}

func (s *PaymentService) UpdateOrder(orderId, status string) error {
	return s.repo.UpdateStatus(orderId, status)
}
