package storage

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/jmoiron/sqlx"
)

type SessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Create() (string, error) {
	b :=make([]byte, 32)
	rand.Read(b)
	token := hex.EncodeToString(b)

	_, err := r.db.Exec("INSERT INTO sessions {token} VALUES ($1)", token)
	return token, err
}

func (r *SessionRepository) Exists(token string) (bool, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM sessions WHERE token = ($1)", token)
	return count > 0, err
}

func(r *SessionRepository) Delete(token string) error {
	_, err := r.db.Exec("DELETE FROM sessions WHERE token =($1)", token)
	return err
} 