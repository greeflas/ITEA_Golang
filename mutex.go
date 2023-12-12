package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	values map[string]int
	m      sync.Mutex
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		values: make(map[string]int),
	}
}

func (c *SafeCounter) Inc(key string) {
	c.m.Lock()
	c.values[key]++
	c.m.Unlock()
}

func (c *SafeCounter) Get(key string) int {
	return c.values[key]
}

func main() {
	var wg sync.WaitGroup

	counter := NewSafeCounter()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc("likes")
		}()
	}

	wg.Wait()

	fmt.Println(counter.Get("likes"))
}
