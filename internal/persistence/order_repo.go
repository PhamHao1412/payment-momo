package persistence

import "payment-momo/internal/entity"

type OrderRepo interface {
	Save(o *entity.Order) error
	UpdateStatus(orderID string, momoResultCode string) error
	Get(orderID string) (*entity.Order, error)
}

type orderRepoInMemory struct{ data map[string]*entity.Order }

func NewOrderRepoInMemory() OrderRepo { return &orderRepoInMemory{data: map[string]*entity.Order{}} }

func (r *orderRepoInMemory) Save(o *entity.Order) error { r.data[o.ID] = o; return nil }

func (r *orderRepoInMemory) UpdateStatus(orderID string, momoResultCode string) error {
	o := r.data[orderID]
	if o == nil {
		return nil
	}
	if momoResultCode == "0" {
		o.Status = "PAID"
	} else {
		o.Status = "FAILED"
	}
	return nil
}

func (r *orderRepoInMemory) Get(orderID string) (*entity.Order, error) { return r.data[orderID], nil }
