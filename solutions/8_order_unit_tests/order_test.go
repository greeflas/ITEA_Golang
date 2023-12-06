package main

import (
	"errors"
	"testing"
	"time"
)

func TestNewOrder(t *testing.T) {
	customer := NewCustomer("tester@example.com")
	creationDate := time.Date(
		2023,
		12,
		1,
		11,
		30,
		45,
		0,
		time.UTC,
	)

	order := NewOrder(customer, creationDate)

	if order.status != OrderStatusInitiated {
		t.Errorf("invalid order status: got: %s, want: %s", order.status, OrderStatusInitiated)
	}

	if !order.createdAt.Equal(creationDate) {
		t.Errorf("invalid createdAt: got: %s, want: %s", order.createdAt, creationDate)
	}

	if !order.updatedAt.Equal(order.createdAt) {
		t.Errorf("updatedAt must be equal createdAt: got: %s, want: %s", order.updatedAt, order.createdAt)
	}
}

func TestOrder_ProcessSuccessfully(t *testing.T) {
	customer := NewCustomer("tester@example.com")
	creationDate := time.Date(
		2023,
		12,
		1,
		11,
		30,
		45,
		0,
		time.UTC,
	)
	processAt := creationDate.Add(time.Hour)

	order := NewOrder(customer, creationDate)
	err := order.Process(processAt)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if order.status != OrderStatusProcessing {
		t.Errorf("invalid order status: got: %s, want: %s", order.status, OrderStatusProcessing)
	}

	if !order.createdAt.Equal(creationDate) {
		t.Errorf("invalid createdAt: got: %s, want: %s", order.createdAt, creationDate)
	}

	if !order.updatedAt.Equal(processAt) {
		t.Errorf("invalid updatedAt: got: %s, want: %s", order.updatedAt, processAt)
	}
}

func TestOrder_ProcessFromIncorrectStatus(t *testing.T) {
	testCases := map[string]func(c *Customer, creatingDate time.Time) *Order{
		"process_to_process": func(c *Customer, d time.Time) *Order {
			order := NewOrder(c, d)
			order.Process(d)

			return order
		},
		"success_to_process": func(c *Customer, d time.Time) *Order {
			order := NewOrder(c, d)
			order.Process(d.Add(time.Hour))
			order.Success(d.Add(time.Hour))

			return order
		},
		"fail_to_process": func(c *Customer, d time.Time) *Order {
			order := NewOrder(c, d)
			order.Process(d.Add(time.Hour))
			order.Fail(d.Add(time.Hour))

			return order
		},
	}

	customer := NewCustomer("tester@example.com")
	creationDate := time.Date(
		2023,
		12,
		1,
		11,
		30,
		45,
		0,
		time.UTC,
	)
	processAt := creationDate.Add(time.Hour)

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			order := tc(customer, creationDate)
			err := order.Process(processAt)

			if !errors.Is(err, ErrInvalidStatusTransition) {
				t.Errorf("invalid err: got: %v, want: %v", err, ErrInvalidStatusTransition)
			}
		})
	}
}
