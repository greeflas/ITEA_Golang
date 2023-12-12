package main

import (
	"fmt"
	"sync"
)

var i int

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc()
		}()
	}

	wg.Wait()

	fmt.Println(i)
}

func inc() {
	i = i + 1
}
