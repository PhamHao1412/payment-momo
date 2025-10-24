package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          string
	Amount      int64
	Description string
	Status      string // PENDING|PAID|FAILED
	CreatedAt   time.Time
}

func NewOrder(amount int64, desc string) *Order {
	return &Order{
		ID:          uuid.NewString(),
		Amount:      amount,
		Description: desc,
		Status:      "PENDING",
		CreatedAt:   time.Now(),
	}
}
