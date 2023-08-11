package util

import (
	"context"
	"crypto/rand"
	"net/http"
	"strings"

	"github.com/geekcamp-vol11-team30/backend/appcontext"
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

func GenerateULID(ctx context.Context) ulid.ULID {
	actx := appcontext.Extract(ctx)
	id, _ := ulid.New(ulid.Timestamp(actx.Now), rand.Reader)
	return id
}

func ULIDFromString(id string) (ulid.ULID, error) {
	id = strings.ToUpper(id)
	ulid, err := ulid.Parse(id)
	if err != nil {
		return ulid, err
	}

	return ulid, nil
}

func ULIDToString(id ulid.ULID) string {
	return strings.ToLower(id.String())
}

func JSONResponse(c echo.Context, code int, data any) error {
	return c.JSON(code, echo.Map{
		"statusCode": code,
		"data":       data,
	})
}
func SetTokenCookie(c echo.Context, cfg config.Config, token entity.Token) {
	c.SetCookie(&http.Cookie{
		Name:  "accessToken",
		Value: token.AccessToken,
		// Path:       "",
		// Domain:     "",
		Expires: token.AccessTokenExpiredAt,
		// RawExpires: "",
		// MaxAge:     0,
		Secure:   cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		// Raw:        "",
		// Unparsed:   []string{},
	})
	c.SetCookie(&http.Cookie{
		Name:  "refreshToken",
		Value: token.RefreshToken,
		// Path:       "",
		// Domain:     "",
		Expires: token.RefreshTokenExpiredAt,
		// RawExpires: "",
		// MaxAge:     0,
		Secure:   cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		// Raw:        "",
		// Unparsed:   []string{},
	})

}
