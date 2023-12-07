package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello!")); err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}

		if _, err := fmt.Fprintf(w, "Hello, %s!", name); err != nil {
			log.Println(err)
		}
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server...")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
