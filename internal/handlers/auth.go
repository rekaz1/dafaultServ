package handlers

import (
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	passwordHash, err := hashPassword(password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = h.Store.CreateUser(r.Context(), email, passwordHash)
	if err != nil {
		http.Error(w, "User already exists or database error", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	user, err := h.Store.GetUserByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if !checkPassword(password, user.PasswordHash) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func hashPassword(password string) (string, error) {
	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash, err := pbkdf2.Key(sha256.New, password, salt, 100000, 32)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x:%x", salt, hash), nil
}

func checkPassword(password string, passwordHash string) bool {
	parts := strings.Split(passwordHash, ":")
	if len(parts) != 2 {
		return false
	}

	salt, err := hex.DecodeString(parts[0])
	if err != nil {
		return false
	}

	expectedHash, err := hex.DecodeString(parts[1])
	if err != nil {
		return false
	}

	actualHash, err := pbkdf2.Key(sha256.New, password, salt, 100000, 32)
	if err != nil {
		return false
	}

	return subtle.ConstantTimeCompare(actualHash, expectedHash) == 1
}
