package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on port 8082")

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("<h1>Hello, World!</h1>"))
		if err != nil {
			return
		}
	})

	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
