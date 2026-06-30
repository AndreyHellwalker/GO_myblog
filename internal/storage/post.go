package storage

import (
	"github.com/AndreyHellwalker/GO_myblog/internal/model"
	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) GetAll() ([]model.Post, error) {
	posts := []model.Post{}
	err := r.db.Select(&posts, "SELECT * FROM posts ORDER BY created_at DESC")
	return posts, err
}

func (r *PostRepository) GetById(id int) (model.Post, error) {
	var post model.Post
	err := r.db.Get(&post, "SELECT * FROM posts WHERE id = $1", id)
	return post, err
}

func (r *PostRepository) Create(post model.Post) error {
	_, err := r.db.Exec(
		"INSERT INTO posts (title, content, image_path) VALUES ($1, $2, $3)",
		post.Title, post.Content, post.ImagePath,
	)
	return err
}