package entity

import "time"

type Token struct {
	AccessToken           string    `json:"accessToken"`
	RefreshToken          string    `json:"refreshToken"`
	AccessTokenExpiredAt  time.Time `json:"accessTokenExpiredAt"`
	RefreshTokenExpiredAt time.Time `json:"refreshTokenExpiredAt"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
