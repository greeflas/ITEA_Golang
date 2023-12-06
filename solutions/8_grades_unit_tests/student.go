package main

type Student struct {
	firstName string
	lastName  string
}

func NewStudent(firstName, lastName string) *Student {
	return &Student{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (s Student) FullName() string {
	return s.firstName + " " + s.lastName
}
