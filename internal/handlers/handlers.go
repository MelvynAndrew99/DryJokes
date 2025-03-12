package handlers

import (
	"net/http"

	"github.com/MelvynAndrew99/DryJokes/internal/jokes"

	"github.com/go-chi/chi"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	/* Serve Static Http files */
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	/* Routes for Jokes */
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	r.Get("/dad", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(jokes.RandomDadJoke()))
	})
	r.Get("/knock", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(jokes.RandomKnockJoke()))
	})
	r.Get("/oneline", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(jokes.RandomOneLiner()))
	})

	return r
}