package store

import (
	"context"
	crand "crypto/rand"
	"database/sql"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type Token string
type UserID string
type User struct {
	ID        UserID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser() (User, error) {
	now := time.Now()
	entropy := rand.New(rand.NewSource(now.UnixNano()))
	ms := ulid.Timestamp(now)
	user := User{
		ID:        UserID(ulid.MustNew(ms, entropy).String()),
		Name:      "no name",
		CreatedAt: now,
		UpdatedAt: now,
	}

	return user, nil
}
func (u *User) Save(ctx context.Context, tx *sql.Tx) error {
	u.UpdatedAt = time.Now()
	_, err := tx.ExecContext(ctx, "INSERT INTO users (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)", u.ID, u.Name, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) RegisterToken(ctx context.Context, tx *sql.Tx, token Token) error {
	now := time.Now()
	_, err := tx.ExecContext(ctx, "INSERT INTO tokens (user_id, id, created_at, updated_at) VALUES (?, ?, ?, ?)", u.ID, token, now, now)
	if err != nil {
		return err
	}
	return nil
}

func NewToken() (Token, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	_, err := crand.Read(b)
	if err != nil {
		return "", err
	}
	for i, v := range b {
		b[i] = letters[v%byte(len(letters))]
	}
	return Token(b), nil
}
