package main

import (
	"context"
	"log"
	"net/http"
	"simpleserver/internal/handlers"
	"simpleserver/internal/storage"
)

func main() {
	ctx := context.Background()

	store, err := storage.New(ctx, "postgres://simpleserver:simpleserver@localhost:5432/simpleserver")
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	mux := http.NewServeMux()

	h := &handlers.Handler{
		Store: store,
	}
	h.Register(mux)

	log.Println("server started at http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
