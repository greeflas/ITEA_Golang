package main

import "fmt"

func main() {
	animals := []string{"raccoon", "cat", "dog"}

	fmt.Println(animals)

	fmt.Println(animals[1])

	animals[0] = "panda"
	fmt.Println(animals)

	fmt.Println("Len:", len(animals))
}
