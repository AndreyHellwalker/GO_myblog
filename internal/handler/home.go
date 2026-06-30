package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my blog")
}

func PostShow(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") 
	fmt.Fprintf(w, "Post ID %s", id)
}