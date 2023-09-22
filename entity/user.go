package entity

import (
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
)

type User struct {
	ID           ulid.ULID `json:"id"`
	Name         string    `json:"name"`
	IsRegistered bool      `json:"isRegistered"`
}

type Calendar struct {
	Events       CalendarEvents `json:"events"`
	Provider     string         `json:"provider"`
	CalendarName string         `json:"calendarName"`
	CalendarID   string         `json:"calendarId"`
	Count        int            `json:"count"`
}
type CalendarEvent struct {
	Name        string     `json:"name"`
	StartTime   *time.Time `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	StartDate   *Date      `json:"startDate"`
	EndDate     *Date      `json:"endDate"`
	IsAllDay    bool       `json:"isAllDay"`
	URL         string     `json:"url"`
	DisplayOnly bool       `json:"displayOnly"`
}
type CalendarEvents []CalendarEvent
type Date struct {
	time.Time
}

const dateFormat = "2006-01-02"

// MarshalJSON implements the json.Marshaler interface.
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format(dateFormat))
}
