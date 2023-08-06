package entity

import (
	"github.com/oklog/ulid/v2"
)

type User struct {
	ID           ulid.ULID `json:"id"`
	Name         string    `json:"name"`
	IsRegistered bool      `json:"isRegistered"`
}

type UserResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	IsRegistered bool   `json:"isRegistered"`
}
