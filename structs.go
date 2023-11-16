package main

import "fmt"

type User struct {
	email string
	age   int
}

// constructor
func NewUser(email string, age int) User {
	return User{
		email: email,
		age:   age,
	}
}

func main() {
	john := User{
		email: "john@example.com",
		age:   23,
	}
	fmt.Printf("%#v\n", john)

	julia := User{
		"julia@example.com",
		30,
	}
	fmt.Printf("%#v\n", julia)

	jane := User{
		email: "jane@example.com",
	}
	fmt.Printf("%#v\n", jane)

	bill := NewUser("bill@example.com", 54)
	fmt.Printf("%#v\n", bill)

	// zero value
	var tester User
	tester = User{
		email: "john@example.com",
		age:   23,
	}
	fmt.Printf("%#v\n", tester)

	fmt.Printf("%v, john: %p tester: %p\n", john == tester, &john, &tester)

	fmt.Println()
	printUserInfo(john)

	fmt.Println()
	incrementAge(&john)
	fmt.Println("main:", john.age)

	fmt.Println()
	cat := struct {
		name   string
		color  string
		weight int
	}{
		name:   "black",
		color:  "red",
		weight: 20,
	}
	fmt.Printf("%#v\n", cat)

	fmt.Println()
	ross := new(User)
	fmt.Println(ross)
	ross.email = "ross@example.com"
	ross.age = 44
	fmt.Println(ross)
}

func printUserInfo(u User) {
	fmt.Printf("E-mail: %s\nAge: %d\n%p\n", u.email, u.age, &u)
}

func incrementAge(u *User) {
	//u.age = u.age + 1
	u.age++

	fmt.Println("incrementAge:", u.age)
}
