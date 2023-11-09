package main

import (
	"fmt"
	"time"
)

func main() {
	inputStatus := 2
	var status string

	switch inputStatus {
	case 1:
		status = "Initiated"
	case 2:
		status = "Approved"
	case 3:
		status = "Declined"
	}

	//if inputStatus == 1 {
	//	status = "Initiated"
	//} else if inputStatus == 2 {
	//	status = "Approved"
	//} else if inputStatus == 3 {
	//	status = "Declined"
	//}

	fmt.Printf("%d is %s\n", inputStatus, status)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("I can rest now.")
	default:
		fmt.Println("Oh, I need to work...")
	}

	switch time.Now().Weekday() {
	case time.Saturday:
		fallthrough
	case time.Sunday:
		fmt.Println("I can rest now.")
	default:
		fmt.Println("Oh, I need to work...")
	}

	a, b := 4, 8
	switch {
	case a > b:
		fmt.Println("a > b")
	case a < b:
		fmt.Println("a < b")
	case a == b:
		fmt.Println("a == b")
	}
}
