//go:generate mockgen -source=./oauth.go -destination=./mock/oauth.go -package=mockcontroller
package controller

import (
	"net/http"
	"time"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type OauthController interface {
	// 各サービスの認証ページにリダイレクト
	RedirectToAuthPage(c echo.Context) error
	// callback
	Callback(c echo.Context) error
	// // 未登録ユーザー作成・トークン発行
	// CreateUnregisteredUserAndToken(c echo.Context) error
	// // トークン更新
	// RefreshToken(c echo.Context) error
}

type oauthController struct {
	cfg       *config.Config
	googleCfg *oauth2.Config
	// uu  usecase.UserUsecase
	// au  usecase.AuthUsecase
}

func NewOauthController(cfg *config.Config, uu usecase.UserUsecase, au usecase.AuthUsecase) OauthController {
	gcfg := &oauth2.Config{
		ClientID:     cfg.OAuth.Google.ClientID,
		ClientSecret: cfg.OAuth.Google.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/oauth2/google/callback",
		Scopes:       []string{"openid", "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/calendar.readonly"},
	}
	return &oauthController{
		cfg:       cfg,
		googleCfg: gcfg,
		// uu:  uu,
		// au:  au,
	}
}

// RedirectToAuthPage implements OauthController.
func (oc *oauthController) RedirectToAuthPage(c echo.Context) error {
	state, err := util.MakeRandomStr(32)
	if err != nil {
		return err
	}
	c.SetCookie(&http.Cookie{
		Name:     "state",
		Value:    state,
		Secure:   oc.cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
	url := oc.googleCfg.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return c.Redirect(302, url)
}

// Callback implements OauthController.
func (oc *oauthController) Callback(c echo.Context) error {
	ctx := c.Request().Context()
	httpClient, _ := oc.googleCfg.Exchange(ctx, c.QueryParam("code"), oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	if httpClient == nil {
		return c.JSON(500, "error1")
	}
	client := oc.googleCfg.Client(ctx, httpClient)

	// service, err := v2.New(client)
	service, err := v2.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return c.JSON(500, "error2")
	}
	// userinfo, err := service.Userinfo.Get().Do()
	userInfo, err := service.Tokeninfo().AccessToken(httpClient.AccessToken).Context(ctx).Do()
	if err != nil {
		return c.JSON(500, "error3")
	}
	// get calendar info
	calendarService, err := calendar.NewService(ctx, option.WithHTTPClient(client), option.WithScopes(calendar.CalendarSettingsReadonlyScope))
	if err != nil {
		return c.JSON(500, "error4")
	}
	calendarList, err := calendarService.CalendarList.List().Context(ctx).Do()
	if err != nil {
		return err
	}
	timeMin := time.Now().Add(-time.Hour * 24 * 7).Format(time.RFC3339)
	timeMax := time.Now().Add(time.Hour * 24 * 365).Format(time.RFC3339)
	events, err := calendarService.Events.List("primary").SingleEvents(true).OrderBy("startTime").TimeMin(timeMin).TimeMax(timeMax).MaxResults(2500).Context(ctx).Do()
	if err != nil {
		return err
		// return c.JSON(500, "error5")
	}
	// events

	return c.JSON(200, map[string]interface{}{
		"userinfo": userInfo,
		"token":    httpClient.AccessToken,
		"client":   httpClient,

		"calendarList": calendarList,
		"events":       events,
	})
}

// // Slash implements SlackController.
// func (ac *oauthController) CreateUnregisteredUserAndToken(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	user, err := ac.uu.CreateAnonymousUser(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	token, err := ac.au.CreateToken(ctx, user)
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(200, token)
// 	// panic("unimplemented")
// }

// // RefreshToken implements OauthController.
// func (ac *oauthController) RefreshToken(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	rtokenReq := entity.RefreshTokenRequest{}
// 	if err := c.Bind(&rtokenReq); err != nil {
// 		return err
// 	}
// 	token, err := ac.au.RefreshToken(ctx, rtokenReq.RefreshToken)
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(200, token)
// }
