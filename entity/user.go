package entity

import (
	"time"

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

type CalendarEvents struct {
	Events []CalendarEvent `json:"events"`
}
type CalendarEvent struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
