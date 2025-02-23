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
	"github.com/jioo/sveltekit-go-chi/api/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var form entity.Login
	var user entity.User
	var validate = validator.New()

	// Decode JSON request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusInternalServerError)
		return
	}

	if err := validate.Struct(form); err != nil {
		utils.ListErrors(w, err)
		return
	}

	ctx := r.Context()
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "unable to get context", http.StatusInternalServerError)
		return
	}

	row := db.QueryRow("SELECT Username, Password FROM users WHERE username = ?", form.Username)
	err = row.Scan(&user.Username, &user.Password)
	if err != nil {
		utils.CustomError(w, "Invalid Username/Password")
		return
	}

	// Verify hased password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		utils.CustomError(w, "Invalid Username/Password")
		return
	}

	token, err := CreateJWT(user.Username)
	if err != nil {
		utils.CustomError(w, fmt.Sprintf("token error: %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	var userID int
	var validate = validator.New()

	// Decode JSON request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusInternalServerError)
		return
	}

	if err := validate.Struct(user); err != nil {
		utils.ListErrors(w, err)
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
		utils.CustomError(w, "Username already exists")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Unable to hash password", http.StatusInternalServerError)
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

func CreateJWT(username string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_KEY"))

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

func VerifyJWT(tokenString string) error {
	secretKey := []byte(os.Getenv("JWT_KEY"))

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
