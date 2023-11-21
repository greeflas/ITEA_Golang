package main

import (
	"fmt"
	"os"
)

func main() {
	homePath := os.Getenv("HOME")
	fmt.Println(homePath)

	os.Setenv("FOO", "bar")
	fooValue := os.Getenv("FOO")
	fmt.Println(fooValue)

	barValue := os.Getenv("BAR")
	fmt.Println(barValue)
}
