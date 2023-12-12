package main

type TransactionType string

const (
	TransactionTypeAuth   TransactionType = "auth"
	TransactionTypeSettle TransactionType = "settle"
	TransactionTypeRefund TransactionType = "refund"
)

type Order struct {
	Id           string        `json:"id"`
	Amount       int           `json:"amount"`
	Transactions []Transaction `json:"transactions"`
}

func (o *Order) HasRefund() bool {
	for _, transaction := range o.Transactions {
		if transaction.Type == TransactionTypeRefund {
			return true
		}
	}

	return false
}

type Transaction struct {
	Id   string          `json:"id"`
	Type TransactionType `json:"type"`
}

func FindOrderWithRefund(orders []*Order) *Order {
	for _, order := range orders {
		if order.HasRefund() {
			return order
		}
	}

	return nil
}
