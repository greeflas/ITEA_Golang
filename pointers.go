package main

import "fmt"

func main() {
	var myNumber int = 8
	fmt.Println(myNumber, &myNumber)

	var ptr *int = &myNumber
	fmt.Println(ptr)

	fmt.Println(*ptr)

	myNumber--
	fmt.Println(myNumber, *ptr)

	*ptr--
	fmt.Println(myNumber, *ptr)

	boolPtr := new(bool)
	fmt.Println(boolPtr, *boolPtr)

	var strPtr *string
	fmt.Println(strPtr)

	if strPtr != nil {
		fmt.Println(*strPtr)
	}

	p1 := &myNumber
	p2 := &myNumber

	fmt.Println(p1 == p2)

	mySecondNumber := 6
	p2 = &mySecondNumber

	fmt.Println(p1 == p2)

	fmt.Println(ptr)
	ptr = nil
	fmt.Println(ptr)

	// part 2
	a := 4

	multiply(a)
	fmt.Println(a, &a)

	multiply2(&a)
	fmt.Println(&a, a)
}

func multiply(num int) {
	num *= 2

	fmt.Println(num, &num)
}

func multiply2(num *int) {
	*num *= 2

	fmt.Println(num, *num)
}
