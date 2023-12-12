package main

import (
	"fmt"
	"math/rand"
)

func main() {
	in := make(chan int)

	go provideNumbers(6, in)

	n, ok := <-in
	fmt.Println(n, ok)

	for num := range in {
		fmt.Println(num)
	}

	n, ok = <-in
	fmt.Println(n, ok)
}

func provideNumbers(max int, out chan int) {
	for i := 0; i < max; i++ {
		out <- rand.Intn(100)
	}
	close(out)
}
