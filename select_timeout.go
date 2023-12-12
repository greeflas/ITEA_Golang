package main

import (
	"fmt"
	"time"
)

func main() {
	consoleIn := make(chan string)

	go readConsole(consoleIn)

	select {
	case data := <-consoleIn:
		fmt.Println("console:", data)
	case <-time.After(time.Second * 3):
		fmt.Println("Time out!")
	}
}

func readConsole(out chan string) {
	fmt.Print("> ")
	var line string
	fmt.Scanln(&line)

	out <- line
}
