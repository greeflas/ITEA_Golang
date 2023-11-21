package main

import "fmt"

func main() {
	//implicit()

	recoverAfterPanicExample()

	fmt.Println("Recovered after panic!!!")
}

func recoverAfterPanicExample() {
	defer func() {
		fmt.Println("Hello from defer!")

		if err := recover(); err != nil {
			fmt.Printf("Recovered after panic! Err: %s\n", err)
		}
	}()

	explicit()

	fmt.Println("After panic!")
}

func explicit() {
	panic("fatal error!")
}

func implicit() {
	//sli := make([]int, 0)
	//_ = sli[100]

	var ptr *string
	*ptr = "pointer"
}
