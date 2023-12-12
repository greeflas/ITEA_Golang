package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("first")

	fmt.Println("second")

	go third()

	go countToTen()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(i + 10)
	}

	time.Sleep(time.Second)
}

func third() {
	fmt.Println("third")
}

func countToTen() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(i)
	}
}
