package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jioo/sveltekit-go-chi/api/db"
	"github.com/jioo/sveltekit-go-chi/api/service"

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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	r.Route("/api/albums", func(r chi.Router) {
		r.Use(withAuth)
		r.Get("/", service.GetAlbums)
		r.Post("/", service.AddAlbum)

		// Subrouters:
		r.Route("/{albumID}", func(r chi.Router) {
			r.Get("/", service.GetAlbumByID)
			r.Put("/", service.UpdateAlbum)
			r.Delete("/", service.DeleteAlbum)
		})
	})

	r.Post("/api/register", service.Register)
	r.Post("/api/login", service.Login)

	// start server
	http.ListenAndServe(":5000", r)
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

func withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		const expected = "Bearer "
		if !strings.HasPrefix(authHeader, expected) {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}
		token := authHeader[len(expected):]

		if err := service.VerifyJWT(token); err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
