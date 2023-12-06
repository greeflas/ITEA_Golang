package main

type Customer struct {
	email string
}

func NewCustomer(email string) *Customer {
	return &Customer{email: email}
}
