package main

import "fmt"

func main() {
	animals := []string{"raccoon", "cat", "dog"}

	fmt.Println(animals)

	fmt.Println(animals[1])

	animals[0] = "panda"
	fmt.Println(animals)

	fmt.Println("Len:", len(animals))

	// part 2
	var sli []bool
	fmt.Println(sli)
	fmt.Println(sli == nil)

	animalsCopy := make([]string, len(animals))
	fmt.Println(animalsCopy)
	fmt.Println(animalsCopy == nil)

	copy(animalsCopy, animals)
	fmt.Println(animalsCopy)

	animalsCopy = append(animalsCopy, "fish")
	fmt.Println(animalsCopy)

	fmt.Println("Len:", len(animalsCopy))
	fmt.Println("Cap:", cap(animalsCopy))

	fmt.Println()

	test := []int{10, 20}
	fmt.Println("Test:", test)
	fmt.Println("Test Len:", len(test))
	fmt.Println("Test Cap:", cap(test))

	test = append(test, 30)
	fmt.Println("Test:", test)
	fmt.Println("Test Len:", len(test))
	fmt.Println("Test Cap:", cap(test))

	fmt.Println()

	homeAnimals := animals[1:]
	fmt.Println(animals)
	fmt.Println(homeAnimals)

	fmt.Println()

	animals[2] = "snake"
	fmt.Println(animals)
	fmt.Println(homeAnimals)

	// part 3
	fmt.Println()
	sliceMutationCase()
}

func sliceMutationCase() {
	sli := make([]int, 0, 3)
	sli = append(sli, 10)
	sli = append(sli, 40)

	fmt.Printf("before\t\tlen(%d) cap(%d) %p %v\n", len(sli), cap(sli), sli, sli)

	mutatedSli := mutateSlice(sli)
	fmt.Printf("after\t\tlen(%d) cap(%d) %p %v\n", len(sli), cap(sli), sli, sli)

	fmt.Printf("mutatedSli\tlen(%d) cap(%d) %p %v\n", len(mutatedSli), cap(mutatedSli), mutatedSli, mutatedSli)
}

func mutateSlice(s []int) []int {
	s = append(s, 90)
	fmt.Printf("mutateSlice\tlen(%d) cap(%d) %p %v\n", len(s), cap(s), s, s)

	return s
}
