package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"`
}

func main() {
	fruits := []string{"apple", "banana", "orange"}

	jsonBytes, err := json.Marshal(fruits)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(string(jsonBytes))

	person := map[string]string{
		"name":  "John",
		"email": "john@example.com",
	}

	jsonBytes, err = json.Marshal(person)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(string(jsonBytes))

	u := user{
		Name:     "Jane",
		Email:    "jane@example.com",
		Password: "qwe",
	}

	jsonBytes, err = json.Marshal(u)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(string(jsonBytes))
}
