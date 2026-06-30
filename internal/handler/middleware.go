package handler

import (
	"net/http"

	"github.com/AndreyHellwalker/GO_myblog/internal/storage"
)

func AuthMiddleware(session *storage.SessionRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			exists, err := session.Exists(cookie.Value)
			if err != nil || !exists {
				http.Redirect(w, r, "/login", http.StatusFound)
				return 
			}

			next.ServeHTTP(w, r)
		})
	}
}