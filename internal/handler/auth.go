package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AndreyHellwalker/GO_myblog/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	sessions *storage.SessionRepository
}

func NewAuthHandler(sessions *storage.SessionRepository) *AuthHandler {
	return &AuthHandler{sessions: sessions}
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Login page — coming soon")
}

func(h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	err := bcrypt.CompareHashAndPassword([]byte(adminPassword), []byte(password))
	if err != nil {
		writeError(w, "invalid password", http.StatusUnauthorized)
		return
	}

	token, err := h.sessions.Create()
	if err != nil {
		writeError(w, "session error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "session",
		Value: token,
		HttpOnly: true,
		Path: "/",
		Expires: time.Now().Add(7 * 24 * time.Hour),
	})

	writeJSON(w, map[string]string{"message": "ok"}, http.StatusOK)
}

func(h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == nil {
		h.sessions.Delete(cookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name: "session",
		Value: "",
		Path: "/",
		Expires: time.Now().Add(7 * 24 * time.Hour),
	})

	writeJSON(w, map[string]string{"meassge": "logged out"}, http.StatusOK)
}