package main

import (
	"testing"
)

func TestFindOrderWithRefund_TransactionsWithRefund(t *testing.T) {
	orders := []*Order{
		{
			Id:     "f9c81316-0bad-4f7c-93df-dd441c5371f2",
			Amount: 1099,
			Transactions: []Transaction{
				{
					Id:   "43c2f68e-85aa-4e1f-a22c-7e42d27a560a",
					Type: TransactionTypeAuth,
				},
				{
					Id:   "2025c1f3-a97a-4f0d-bc2c-dcbcea63930a",
					Type: TransactionTypeSettle,
				},
			},
		},
		{
			Id:     "8e08894d-0c8b-475c-8686-5ed147cb13f0",
			Amount: 300,
			Transactions: []Transaction{
				{
					Id:   "7ee3f3e3-de15-4f43-827e-802a5376953f",
					Type: TransactionTypeAuth,
				},
				{
					Id:   "86ae4de9-55d0-4132-b541-fe3e33c6f838",
					Type: TransactionTypeRefund,
				},
			},
		},
	}

	order := FindOrderWithRefund(orders)
	expectedOrderId := "8e08894d-0c8b-475c-8686-5ed147cb13f0"

	if order == nil {
		t.Errorf("got: %v, want: %s", order, expectedOrderId)
	}

	if order.Id != expectedOrderId {
		t.Errorf("invalid order ID: got: %s, want: %s", order.Id, expectedOrderId)
	}
}

func TestFindOrderWithRefund_TransactionsWithoutRefund(t *testing.T) {
	orders := []*Order{
		{
			Id:     "f9c81316-0bad-4f7c-93df-dd441c5371f2",
			Amount: 1099,
			Transactions: []Transaction{
				{
					Id:   "43c2f68e-85aa-4e1f-a22c-7e42d27a560a",
					Type: TransactionTypeAuth,
				},
				{
					Id:   "2025c1f3-a97a-4f0d-bc2c-dcbcea63930a",
					Type: TransactionTypeSettle,
				},
			},
		},
		{
			Id:     "8e08894d-0c8b-475c-8686-5ed147cb13f0",
			Amount: 300,
			Transactions: []Transaction{
				{
					Id:   "7ee3f3e3-de15-4f43-827e-802a5376953f",
					Type: TransactionTypeAuth,
				},
				{
					Id:   "86ae4de9-55d0-4132-b541-fe3e33c6f838",
					Type: TransactionTypeSettle,
				},
			},
		},
	}

	order := FindOrderWithRefund(orders)

	if order != nil {
		t.Errorf("got: %s, want: %v", order.Id, nil)
	}
}
