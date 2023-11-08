package main

import "fmt"

type Status int

//const (
//	Initial Status = 1
//	Processing Status = 2
//	Success Status = 3
//	Fail Status = 4
//)

const (
	Initial Status = iota + 1
	Processing
	Success
	Fail
)

func main() {
	var st Status
	st = Initial

	fmt.Println("Status:", st)

	fmt.Println(Initial, Processing, Success, Fail)
}
