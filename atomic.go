package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var i atomic.Int64

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

	fmt.Println(i.Load())
}

func inc() {
	i.Add(1)
}
