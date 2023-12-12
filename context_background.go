package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	ctx, cancel := context.WithCancel(ctx)

	consoleIn := make(chan string)

	go readConsole(ctx, consoleIn)

	for data := range consoleIn {
		if data == "exit" {
			cancel()
			break
		}

		fmt.Println("console:", data)
	}

	//select {
	//case data := <-consoleIn:
	//	fmt.Println("console:", data)
	//case <-ctx.Done():
	//	fmt.Println("Time out!")
	//	cancel()
	//}
}

func readConsole(ctx context.Context, out chan string) {
	for {
		fmt.Print("> ")
		var line string
		fmt.Scanln(&line)

		select {
		case <-ctx.Done():
			return
		case out <- line:
		}
	}
}
