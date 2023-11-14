package main

import "fmt"

func main() {
	hello := func() {
		fmt.Println("Hello")
	}

	hello()

	pln := fmt.Println
	pln("Test")

	var hi func(string)

	hi = func(message string) {
		fmt.Println(message)
	}

	hi("Hi!")

	sayHiToFiend := createGreeting("Hi, %s! What's up?")

	sayHiToFiend("John")
	sayHiToFiend("Jane")
}

func createGreeting(pattern string) func(string) {
	return func(name string) {
		fmt.Printf(pattern+"\n", name)
	}
}
