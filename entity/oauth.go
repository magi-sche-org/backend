package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type OauthProvider struct {
	ID           ulid.ULID
	Name         string
	ClientId     string
	ClientSecret string
}
type OauthUserInfo struct {
	ID                    ulid.ULID
	UserId                ulid.ULID
	ProviderId            ulid.ULID
	ProviderUid           string
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt *time.Time
}
