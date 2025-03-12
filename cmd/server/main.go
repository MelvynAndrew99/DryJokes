package main

import (
	"log"
	"net/http"

	"github.com/MelvynAndrew99/DryJokes/internal/handlers"
)

func main() {
	r := handlers.SetupRouter()
	log.Println("Starting Server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}