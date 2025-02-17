package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/jioo/sveltekit-go-chi/api/entity"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	var albums []entity.Album

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "context error", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var alb entity.Album
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		albums = append(albums, alb)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	var album entity.Album
	albumID := chi.URLParam(r, "albumID")

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "context error", http.StatusInternalServerError)
		return
	}

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", albumID)
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	if err == sql.ErrNoRows {
		http.Error(w, fmt.Sprintf("album with id %v not found", albumID), http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("error querying album: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

func AddAlbum(w http.ResponseWriter, r *http.Request) {
	var newAlbum entity.Album

	// Decode JSON request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newAlbum)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "unable to get context", http.StatusInternalServerError)
		return
	}

	result, err := db.Exec(
		"INSERT INTO album (title, artist, price) VALUES (?, ?, ?)",
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Price,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("error inserting album: %v", err), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting last insert id: %v", err), http.StatusInternalServerError)
		return
	}
	newAlbum.ID = fmt.Sprintf("%d", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newAlbum)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var album entity.Album
	albumID := chi.URLParam(r, "albumID")

	// Decode JSON request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&album)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "unable to get context", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(
		"UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?",
		album.Title,
		album.Artist,
		album.Price,
		albumID,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating album: %v", err), http.StatusInternalServerError)
		return
	}

	album.ID = albumID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	albumID := chi.URLParam(r, "albumID")

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "unable to get context", http.StatusInternalServerError)
		return
	}

	_, err := db.Exec(
		"DELETE FROM album WHERE ID = ?",
		albumID,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting album: %v, id: %v", err, albumID), http.StatusInternalServerError)
		return
	}
}
