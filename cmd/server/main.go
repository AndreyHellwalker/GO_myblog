package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AndreyHellwalker/GO_myblog/internal/handler"
	"github.com/AndreyHellwalker/GO_myblog/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main () {
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db := storage.NewDB(dsn)
	defer db.Close()

	postRepo := storage.NewPostRepository(db)
	sessionRepo := storage.NewSessionRepository(db)

	postHandler := handler.NewPostHandler(postRepo)
	authHandler := handler.NewAuthHandler(sessionRepo)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Cookie"},
		AllowCredentials: true,
	}))

	r.Get("/", handler.Home)
	r.Post("/login", authHandler.Login)
	r.Post("/logout", authHandler.Logout)

	r.Get("/posts", postHandler.List)
	r.Get("/posts/{id}", postHandler.Show)

	r.Group(func(r chi.Router) {
		r.Use(handler.AuthMiddleware(sessionRepo))
		r.Post("/posts", postHandler.Create)
		r.Post("/upload", handler.UploadImage)
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}