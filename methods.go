package main

import "fmt"

type Customer struct {
	firstName string
	lastName  string
	email     string
}

func (c *Customer) FullName() string {
	return c.firstName + " " + c.lastName
}

func (c *Customer) ChangeName(firstName, lastName string) {
	c.firstName = firstName
	c.lastName = lastName
}

func NewCustomer(firstName string, lastName string, email string) *Customer {
	return &Customer{firstName: firstName, lastName: lastName, email: email}
}

type OrderStatus string

func (s OrderStatus) ToString() string {
	return string(s)
}

func main() {
	customer := NewCustomer("John", "Doe", "john.doe@example.com")
	fmt.Println(customer.FullName())

	customer.ChangeName("Jane", "Doe")
	fmt.Println(customer.FullName())

	os := OrderStatus("initiate")
	fmt.Println(os.ToString())
}
