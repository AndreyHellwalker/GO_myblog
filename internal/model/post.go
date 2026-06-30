package model

import "time"

type Post struct {
	ID int `db:"id"`
	Title string `db:"title"`
	Content string `db:"content"`
	ImagePath string `db:"image_path"`
	CreatedAt time.Time `db:"created_at"`
}