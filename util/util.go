package util

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"net/smtp"
	"strings"
	"time"

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
func GenerateULIDNow() ulid.ULID {
	id, _ := ulid.New(ulid.Timestamp(time.Now()), rand.Reader)
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
		Path:  "/",
		// Domain:     "",
		Expires: token.AccessTokenExpiredAt,
		// RawExpires: "",
		// MaxAge:     0,
		Secure:   cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		// Raw:        "",
		// Unparsed:   []string{},
	})
	c.SetCookie(&http.Cookie{
		Name:  "refreshToken",
		Value: token.RefreshToken,
		Path:  "/",
		// Domain:     "",
		Expires: token.RefreshTokenExpiredAt,
		// RawExpires: "",
		// MaxAge:     0,
		Secure:   cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		// Raw:        "",
		// Unparsed:   []string{},
	})

}
func DeleteTokenCookie(c echo.Context, cfg config.Config) {
	c.SetCookie(&http.Cookie{
		Name:  "accessToken",
		Value: "",
		Path:  "/",
		// Expires:  time.Now(),
		Secure:   cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSite(cfg.TokenSameSite),
		MaxAge:   -1,
	})
	c.SetCookie(&http.Cookie{
		Name:  "refreshToken",
		Value: "",
		Path:  "/",
		// Expires:  time.Now(),
		Secure:   cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSite(cfg.TokenSameSite),
		MaxAge:   -1,
	})
}

func MakeRandomStr(digit int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

func SendMail(cfg config.Config, targetAddrs string, title string, body string) error {
	fmt.Println("send email called!!!!!!", targetAddrs, title, body)
	hostname := cfg.SMTP.Host // SMTPサーバーのホスト名
	port := cfg.SMTP.Port     // SMTPサーバーのポート番号
	password := cfg.SMTP.Password
	// from := cfg.SMTP.Email // 送信者のメールアドレス
	// from = "マジスケ<noreply@magi-sche.net>"
	username := cfg.SMTP.User

	recipients := []string{targetAddrs} // 送信先のメールアドレス
	// auth := smtp.PlainAuth("", username, password, hostname)
	auth := smtp.PlainAuth("", username, password, hostname)

	msg := []byte(strings.ReplaceAll(fmt.Sprintf(
		"From: マジスケ<noreply@magi-sche.net>\nTo: %s\nSubject: %s\n\n%s", strings.Join(recipients, ","), title, body),
		"\n", "\r\n"))
	// メール送信
	err := smtp.SendMail(fmt.Sprintf("%s:%d", hostname, port), auth, "noreply@magi-sche.net", recipients, msg)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil

}
