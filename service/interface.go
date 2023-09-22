package service

import (
	"context"
	"time"

	"github.com/geekcamp-vol11-team30/backend/entity"
	"golang.org/x/oauth2"
)

type OauthCalendarService interface {
	GetAuthURL(ctx context.Context) (url string, state string, err error)
	ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error)
	GetOrCreateUserByCode(ctx context.Context, code string, user *entity.User) (*entity.User, error)
	GetPrimaryCalendar(ctx context.Context, oui entity.OauthUserInfo, timeMin *time.Time, timeMax *time.Time) (entity.Calendar, error)
}
