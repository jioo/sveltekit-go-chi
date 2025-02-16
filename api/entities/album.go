package entities

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	var albums []album

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "context error", http.StatusInternalServerError)
	}

	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	for rows.Next() {
		var alb album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		albums = append(albums, alb)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	var album album
	albumID := chi.URLParam(r, "albumID")

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "context error", http.StatusInternalServerError)
	}

	err := db.QueryRow("SELECT * FROM album WHERE id = ?", albumID).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)

	if err == sql.ErrNoRows {
		http.Error(w, fmt.Sprintf("album with id %v not found", albumID), http.StatusInternalServerError)
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("error querying album: %v", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

func AddAlbum(w http.ResponseWriter, r *http.Request) {
	var newAlbum album

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&newAlbum); err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "unable to get context", http.StatusInternalServerError)
	}

	result, err := db.Exec(
		"INSERT INTO album (title, artist, price) VALUES (?, ?, ?)",
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Price,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("error inserting album: %v", err), http.StatusInternalServerError)
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting last insert id: %v", err), http.StatusInternalServerError)
	}
	newAlbum.ID = fmt.Sprintf("%d", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newAlbum)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var album album

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "unable to get context", http.StatusInternalServerError)
	}

	result, err := db.Exec(
		"UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?",
		album.Title,
		album.Artist,
		album.Price,
		album.ID,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating album: %v", err), http.StatusInternalServerError)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting rows affected: %v", err), http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("album with id %s not found", album.ID), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rowsAffected)
}
