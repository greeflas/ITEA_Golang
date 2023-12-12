package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	//wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("first")
	}()

	fmt.Println("second")

	//wg.Add(1)
	go func() {
		defer wg.Done()
		third()
	}()

	wg.Wait()
}

func third() {
	fmt.Println("third")
}
