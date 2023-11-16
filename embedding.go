package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (p Person) FullName() string {
	return p.firstName + " " + p.lastName
}

type Employee struct {
	Person
	position string
}

func (e Employee) FullName() string {
	return e.Person.FullName() + " (" + e.position + ")"
}

func main() {
	e := Employee{
		Person: Person{
			firstName: "John",
			lastName:  "Doe",
		},
		position: "Software Engineer",
	}

	fmt.Printf("%#v\n", e)
	fmt.Println(e.FullName())
}
