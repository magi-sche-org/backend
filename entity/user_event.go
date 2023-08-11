package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type UserEventAnswer struct {
	ID           ulid.ULID             `json:"id"`
	UserID       ulid.ULID             `json:"userId"`
	EventID      ulid.ULID             `json:"-"`
	UserNickname string                `json:"userNickname"`
	Note         string                `json:"note"`
	Units        []UserEventAnswerUnit `json:"units"`
}
type UserEventAnswerRequest struct {
	UserNickname string                       `json:"userNickname"`
	Note         string                       `json:"note"`
	Units        []UserEventAnswerUnitRequest `json:"units"`
}
type UserEventAnswerResponse struct {
	ID           string                        `json:"id"`
	UserID       string                        `json:"userId"`
	IsYourAnswer bool                          `json:"isYourAnswer"`
	UserNickname string                        `json:"userNickname"`
	Note         string                        `json:"note"`
	Units        []UserEventAnswerUnitResponse `json:"units"`
}

type Availability string

const (
	Available   Availability = "available"
	Maybe       Availability = "maybe"
	Unavailable Availability = "unavailable"
	Error       Availability = "error"
)

type UserEventAnswerUnit struct {
	ID                ulid.ULID    `json:"id"`
	UserEventAnswerID ulid.ULID    `json:"answerId"`
	EventTimeUnitID   ulid.ULID    `json:"eventTimeUnitId"`
	Availability      Availability `json:"availability"`
}
type UserEventAnswerUnitRequest struct {
	EventTimeUnitID string       `json:"eventTimeUnitId"`
	Availability    Availability `json:"availability"`
}
type UserEventAnswerUnitResponse struct {
	EventTimeUnitID string       `json:"eventTimeUnitId"`
	Availability    Availability `json:"availability"`
	StartsAt        time.Time    `json:"startsAt"`
	EndsAt          time.Time    `json:"endsAt"`
}
