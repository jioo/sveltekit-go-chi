package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// setup router and middleware
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// setup routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// start server
	http.ListenAndServe(":8080", r)
}
