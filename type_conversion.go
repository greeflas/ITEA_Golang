package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 6
	var b float64 = float64(5)

	fmt.Println(a, b)

	var c float64 = 4.56
	var d int = int(c)

	fmt.Println(c, d)

	var str string = "Hello"
	var bytes []byte = []byte(str)

	fmt.Println(str, bytes)

	num := "56"
	number, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(number)

	number = -67
	num = strconv.Itoa(number)

	fmt.Println(num)
}
