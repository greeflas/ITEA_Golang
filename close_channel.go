package main

import (
	"fmt"
	"math/rand"
)

func main() {
	in := make(chan int)
	errors := make(chan error)

	go provideNumbers(6, in, errors)

	n, ok := <-in
	fmt.Println(n, ok)

	for num := range in {
		fmt.Println(num)
	}

	n, ok = <-in
	fmt.Println(n, ok)
}

func provideNumbers(max int, out chan int, errors chan error) {
	for i := 0; i < max; i++ {
		out <- rand.Intn(100)
	}
	close(out)
}
