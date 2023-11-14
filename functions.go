package main

import "fmt"

func main() {
	sayHello()
	sayHello()
	sayHello()

	values := getInitialValues()
	fmt.Println(values)

	first := "World"
	second := "Hello"

	first, second = swapByValue(first, second)
	fmt.Println(first, second)

	fmt.Println(sum(2, 4, 4))

	nums := []int{4, 66, 30}
	fmt.Println(sum(nums...))
}

func sayHello() {
	fmt.Println("Hello, World!")
}

func getInitialValues() []int {
	return []int{4, 5}
}

// func swapByValue(a string, b string) {
func swapByValue(a, b string) (string, string) {
	return b, a
}

// variadic function
func sum(numbers ...int) int {
	var res int
	for _, number := range numbers {
		res += number
	}

	return res
}
