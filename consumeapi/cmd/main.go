package main

import (
	"consumeapi/internal/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/posts", handlers.GetAllPosts)
	mux.HandleFunc("POST /api/numbers", handlers.Add)

	log.Println("starting server on :", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
