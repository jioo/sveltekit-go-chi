package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jioo/sveltekit-go-chi/api/entity"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	var userID int

	// Decode JSON request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(user); err != nil {
		http.Error(w, fmt.Sprintf("validation error: %v", err), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "unable to get context", http.StatusInternalServerError)
		return
	}

	row := db.QueryRow("SELECT ID FROM users WHERE username = ?", user.Username)
	err = row.Scan(&userID)
	if err != nil {
		if err != sql.ErrNoRows {
			http.Error(w, fmt.Sprintf("database error: %v", err), http.StatusInternalServerError)
			return
		}
	}
	if userID != 0 {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusInternalServerError)
		return
	}

	result, err := db.Exec(
		"INSERT INTO users (username, password, firstName, lastName) VALUES (?, ?, ?, ?)",
		user.Username,
		hashedPassword,
		user.FirstName,
		user.LastName,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("error inserting user: %v", err), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting last insert id: %v", err), http.StatusInternalServerError)
		return
	}
	user.ID = fmt.Sprintf("%d", id)
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createJWT(username string) (string, error) {
	secretKey := os.Getenv("JWT_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyJWT(tokenString string) error {
	secretKey := os.Getenv("JWT_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
