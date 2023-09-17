//go:generate mockgen -source=./oauth.go -destination=./mock/oauth.go -package=mockcontroller
package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/geekcamp-vol11-team30/backend/appcontext"
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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
	oau       usecase.OauthUsecase
	uu        usecase.UserUsecase
	au        usecase.AuthUsecase
}

func NewOauthController(cfg *config.Config, oau usecase.OauthUsecase, uu usecase.UserUsecase, au usecase.AuthUsecase) OauthController {
	gcfg := &oauth2.Config{
		ClientID:     cfg.OAuth.Google.ClientID,
		ClientSecret: cfg.OAuth.Google.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  fmt.Sprintf("%s/oauth2/google", cfg.BaseURL), // "http://localhost:8080/oauth2/google/callback",
		Scopes:       []string{"openid", "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/calendar.readonly"},
	}
	return &oauthController{
		cfg:       cfg,
		googleCfg: gcfg,
		oau:       oau,
		uu:        uu,
		au:        au,
	}
}

// RedirectToAuthPage implements OauthController.
func (oc *oauthController) RedirectToAuthPage(c echo.Context) error {
	// get next parameter if exists
	next := c.QueryParam("next")
	if next == "" {
		next = oc.cfg.OAuth.DefaultReturnURL
	}

	url, state, err := oc.oau.GetGoogleAuthURL(c.Request().Context())
	if err != nil {
		return err
	}
	c.SetCookie(&http.Cookie{
		Name:     "state",
		Value:    state,
		Secure:   oc.cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	c.SetCookie(&http.Cookie{
		Name:     "next",
		Value:    next,
		Secure:   oc.cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return c.Redirect(302, url)
}

// Callback implements OauthController.
func (oc *oauthController) Callback(c echo.Context) error {
	ctx := c.Request().Context()
	// state check
	stateCookie, err := c.Cookie("state")
	if err != nil {
		return err
	}
	// remove state
	c.SetCookie(&http.Cookie{
		Name:     "state",
		Value:    "",
		Secure:   oc.cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})
	// get next
	nextCookie, err := c.Cookie("next")
	if err != nil {
		return err
	}
	// remove next
	c.SetCookie(&http.Cookie{
		Name:     "next",
		Value:    "",
		Secure:   oc.cfg.Env != "dev",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	if stateCookie.Value != c.QueryParam("state") {
		return c.JSON(400, "invalid state")
	}

	// get token
	token, err := oc.oau.LoginGoogleWithCode(ctx, c.QueryParam("code"))
	if err != nil {
		return err
	}

	acuser, err := appcontext.Extract(ctx).GetUser()
	var acuserp *entity.User
	if err != nil {
		acuserp = nil
	} else {
		acuserp = &acuser
	}
	log.Println(acuser, acuserp)
	user, err := oc.oau.FetchAndRegisterOauthUserInfo(ctx, token, acuserp)
	if err != nil {
		return err
	}
	log.Println(acuser, acuserp, user)
	sToken, err := oc.au.CreateToken(ctx, user)
	if err != nil {
		return err
	}
	util.SetTokenCookie(c, *oc.cfg, sToken)
	// client := oc.googleCfg.Client(ctx, token)

	// service, err := v2.NewService(ctx, option.WithHTTPClient(client))
	// if err != nil {
	// 	return c.JSON(500, "error2")
	// }
	// // userinfo, err := service.Userinfo.Get().Do()
	// userInfo, err := service.Tokeninfo().AccessToken(token.AccessToken).Context(ctx).Do()
	// if err != nil {
	// 	return c.JSON(500, "error3")
	// }
	// get calendar info
	// calendarService, err := calendar.NewService(ctx, option.WithHTTPClient(client), option.WithScopes(calendar.CalendarSettingsReadonlyScope))
	// if err != nil {
	// 	return c.JSON(500, "error4")
	// }
	// calendarList, err := calendarService.CalendarList.List().Context(ctx).Do()
	// if err != nil {
	// 	return err
	// }
	// timeMin := time.Now().Add(-time.Hour * 24 * 7).Format(time.RFC3339)
	// timeMax := time.Now().Add(time.Hour * 24 * 365).Format(time.RFC3339)
	// events, err := calendarService.Events.List("primary").SingleEvents(true).OrderBy("startTime").TimeMin(timeMin).TimeMax(timeMax).MaxResults(2500).Context(ctx).Do()
	// if err != nil {
	// 	return err
	// 	// return c.JSON(500, "error5")
	// }
	// // events
	// return c.Redirect(302, oc.cfg.OAuth.DefaultReturnURL)
	return c.Redirect(302, nextCookie.Value)

	// return c.JSON(200, map[string]interface{}{
	// 	// "userinfo": userInfo,
	// 	"token":  token.AccessToken,
	// 	"client": token,

	// 	// "calendarList": calendarList,
	// 	// "events":       events,
	// })
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
