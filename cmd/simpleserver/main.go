package main

import (
	"log"
	"net/http"

	"simpleserver/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	handlers.Register(mux)

	log.Println("server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
