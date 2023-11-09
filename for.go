package main

import "fmt"

func main() {
	number := 1

	number = number + 1
	number += 1
	number++

	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	i := 0
	for i < 10 {
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()

	for {
		i++
		fmt.Print(i, " ")

		if i == 15 {
			break
		}
	}

	fmt.Println()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Print(i, " ")
	}

	fmt.Println()

outer:
	for i := 0; i < 3; i++ {
		fmt.Println("outer")

		for j := 0; j < 3; j++ {
			fmt.Printf("\tinner\n")

			if i == 1 && j == 1 {
				break outer
			}
		}
	}
}
