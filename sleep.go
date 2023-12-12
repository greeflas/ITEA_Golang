package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	time.Sleep(time.Second * 3)

	fmt.Println(time.Now())
}
