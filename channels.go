package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	fmt.Println(ch)

	//ch = make(chan int, 3) // buffered channel
	ch = make(chan int) // unbuffered channel

	data := [][]int{
		{3, 7, 8},
		{1, 79, 9},
		{7, 3, 5},
	}

	for _, nums := range data {
		go sum(nums, ch)
	}

	//for i := 0; i < len(data); i++ {
	//	fmt.Println(<-ch)
	//}
	time.Sleep(time.Second * 2)
	fmt.Println(<-ch)

	time.Sleep(time.Second * 2)
	fmt.Println(<-ch)

	time.Sleep(time.Second * 2)
	fmt.Println(<-ch)

	time.Sleep(time.Second)
}

func sum(nums []int, out chan int) {
	res := 0

	for _, num := range nums {
		res += num
	}

	out <- res
	fmt.Println("sec:", time.Now().Second())
}
