package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	in := make(chan int)
	errorsCh := make(chan error)

	const maxNumbers = 6

	go provideNumbers(maxNumbers, in, errorsCh)

	n, ok := <-in
	fmt.Println(n, ok)

	for i := 0; i < maxNumbers; i++ {
		select {
		case num := <-in:
			fmt.Println(num)
		case err := <-errorsCh:
			fmt.Println("error:", err)
		}
	}

	n, ok = <-in
	fmt.Println(n, ok)
}

func provideNumbers(max int, out chan int, errorsCh chan error) {
	for i := 0; i < max; i++ {
		number := rand.Intn(100)

		if number < 50 {
			errorsCh <- errors.New("error occurred")
			continue
		}

		out <- number
	}
	close(out)
}
