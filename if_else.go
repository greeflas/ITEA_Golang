package main

import "fmt"

func main() {
	a := 4
	b := 6

	if a < b {
		fmt.Println("a < b")
	}

	if a > b {
		fmt.Println("a < b")
	}

	if a == b {
		fmt.Println("a == b")
	}

	if a != b {
		fmt.Println("a != b")
	}

	if a <= b {
		fmt.Println("a <= b")
	}

	if a >= b {
		fmt.Println("a >= b")
	}

	c := 7
	d := 23

	if a == b && c == d {
		fmt.Println("a == b && c == d")
	}

	if a == b || c == d {
		fmt.Println("a == b || c == d")
	}

	if !(a == b) {
		fmt.Println("!(a == b)")
	}

	const (
		adminLogin = "admin"
		adminPass  = "123"
	)

	const (
		moderatorLogin = "moderator"
		moderatorPass  = "321"
	)

	var login, password string

	fmt.Print("Enter login: ")
	fmt.Scan(&login)

	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	if login == adminLogin && password == adminPass {
		fmt.Println("Hello, Admin! Welcome to admin panel.")
	} else if login == moderatorLogin && password == moderatorPass {
		fmt.Println("Hello, Moderator!")
	} else {
		fmt.Println("Invalid credentials!")
	}
}
