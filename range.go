package main

import "fmt"

func main() {
	languages := []string{"C", "C++", "TypeScript", "PHP", "Go"}

	for i := 0; i < len(languages); i++ {
		fmt.Printf("%d -> %s\n", i, languages[i])
	}

	fmt.Println()

	for i, language := range languages {
		fmt.Printf("%d -> %s\n", i, language)
	}

	fmt.Println()

	for i := range languages {
		fmt.Print(i, " ")
	}

	fmt.Println()

	for _, language := range languages {
		fmt.Print(language, " ")
	}
}
