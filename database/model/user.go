package model

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Email    string
	IsActive bool
}

func NewUser(id uuid.UUID, email string) *User {
	return &User{
		Id:       id,
		Email:    email,
		IsActive: true,
	}
}

func (u User) PrintInfo() {
	fmt.Printf("ID:\t%s\n", u.Id.String())
	fmt.Printf("E-mail:\t%s\n", u.Email)

	if u.IsActive {
		fmt.Println("This is active user.")
	}
}
