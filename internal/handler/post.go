package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/AndreyHellwalker/GO_myblog/internal/model"
	"github.com/AndreyHellwalker/GO_myblog/internal/storage"
	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	repo *storage.PostRepository
}

func NewPostHandler(repo *storage.PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
} 

// @Summary      Список постов
// @Tags         posts
// @Produce      json
// @Success      200  {object}  Response
// @Router       /posts [get]
func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	posts, err := h.repo.GetAll()
	if err != nil {
		writeError(w, "cannot get posts", http.StatusInternalServerError)
		return
	}
	
	writeJSON(w, posts, http.StatusOK)
}

// @Summary      Получить пост
// @Tags         posts
// @Produce      json
// @Param        id   path      int  true  "ID поста"
// @Success      200  {object}  Response
// @Failure      404  {object}  Response
// @Router       /posts/{id} [get]
func (h *PostHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, "invalid id", http.StatusBadRequest)
		return
	}

	post, err := h.repo.GetById(id)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, "post not found", http.StatusNotFound)
	}
	if err != nil {
		writeError(w, "post not found", http.StatusNotFound)
		return
	}

	writeJSON(w, post, http.StatusOK)
}

// @Summary      Создать пост
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post  body      model.Post  true  "Данные поста"
// @Success      201   {object}  Response
// @Router       /posts [post]
func (h* PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		writeError(w, "invalid body", http.StatusBadRequest)
		return
	}

	if post.Title == "" || post.Content == "" {
		writeError(w, "title and content requred", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(post); err != nil {
		writeError(w, "cannot create post", http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{"message": "created"}, http.StatusCreated)
}