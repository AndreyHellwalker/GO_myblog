package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AndreyHellwalker/GO_myblog/internal/model"
)

type Response struct {
	Data any `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

type PostsResponse struct {
	Data  []model.Post `json:"data,omitempty"`
	Error string       `json:"error,omitempty"`
}

type PostResponse struct {
	Data  model.Post `json:"data,omitempty"`
	Error string     `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Data: data})
}

func writeError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Error: message})
}

