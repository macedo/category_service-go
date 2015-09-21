package main

import (
	"log"
	"net/http"
)

const (
	port string = ":8080"
)

func main() {
	router := NewRouter()

	log.Println("Category Service star on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
