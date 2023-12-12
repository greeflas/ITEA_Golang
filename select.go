package main

import (
	"fmt"
	"net/http"
)

func main() {
	consoleIn := make(chan string)
	serverIn := make(chan string)

	go readConsole(consoleIn)
	go startServer(serverIn)

	select {
	case data := <-consoleIn:
		fmt.Println("console:", data)
	case data := <-serverIn:
		fmt.Println("server:", data)
	}
}

func readConsole(out chan string) {
	fmt.Print("> ")
	var line string
	fmt.Scanln(&line)

	out <- line
}

func startServer(out chan string) {
	http.HandleFunc("/in", func(w http.ResponseWriter, r *http.Request) {
		out <- r.URL.Query().Get("data")
	})

	http.ListenAndServe(":8080", nil)
}
