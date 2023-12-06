package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrInvalidStatusTransition = errors.New("invalid status transition")

type OrderStatus string

const (
	OrderStatusInitiated  OrderStatus = "initiated"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusSuccess    OrderStatus = "success"
	OrderStatusFail       OrderStatus = "fail"
)

type Order struct {
	customer  *Customer
	status    OrderStatus
	createdAt time.Time
	updatedAt time.Time
}

func NewOrder(customer *Customer, createdAt time.Time) *Order {
	return &Order{
		customer:  customer,
		status:    OrderStatusInitiated,
		createdAt: createdAt,
		updatedAt: createdAt,
	}
}

func (o *Order) Process(processAt time.Time) error {
	if o.status != OrderStatusInitiated {
		return fmt.Errorf("process from status: %q: %w", o.status, ErrInvalidStatusTransition)
	}

	o.status = OrderStatusProcessing
	o.updatedAt = processAt

	return nil
}

func (o *Order) Success(successAt time.Time) error {
	if o.status != OrderStatusProcessing {
		return fmt.Errorf("success from status: %q: %w", o.status, ErrInvalidStatusTransition)
	}

	o.status = OrderStatusSuccess
	o.updatedAt = successAt

	return nil
}

func (o *Order) Fail(failAt time.Time) error {
	if o.status != OrderStatusProcessing {
		return fmt.Errorf("fail from status: %q: %w", o.status, ErrInvalidStatusTransition)
	}

	o.status = OrderStatusFail
	o.updatedAt = failAt

	return nil
}
