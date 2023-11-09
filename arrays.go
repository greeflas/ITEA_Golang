package main

import "fmt"

func main() {
	var arr [3]int // zero value: [0 0 0]
	fmt.Println(arr)

	//var animals [3]string = [3]string{"raccoon", "cat", "dog"}
	//var animals = [3]string{"raccoon", "cat", "dog"}
	//animals := [3]string{"raccoon", "cat", "dog"}
	animals := [...]string{"raccoon", "cat", "dog"}
	fmt.Println(animals)

	fmt.Println(animals[1])

	animals[0] = "panda"
	fmt.Println(animals)

	fmt.Println("Len:", len(animals))

	animalsCopy := animals
	fmt.Println(animalsCopy)

	animalsCopy[0] = "parrot"
	fmt.Println(animals, animalsCopy)
}
