package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jioo/sveltekit-go-chi/api/db"
	"github.com/jioo/sveltekit-go-chi/api/entities"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// setup router
	r := chi.NewRouter()

	// setup middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// setup db connection
	r.Use(dbContext)

	// setup routes
	r.Route("/albums", func(r chi.Router) {
		r.Get("/", entities.GetAlbums)
		r.Post("/", entities.AddAlbum)

		// Subrouters:
		r.Route("/{albumID}", func(r chi.Router) {
			r.Get("/", entities.GetAlbumByID)
			r.Put("/", entities.UpdateAlbum)
			// r.Delete("/", entities.DeleteAlbum)
		})
	})

	// start server
	http.ListenAndServe(":8080", r)
}

func dbContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := db.Connect()
		if err != nil {
			log.Fatal(err)
		}

		// close the connection after request
		defer db.Close()

		ctx := context.WithValue(r.Context(), "db", db)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
