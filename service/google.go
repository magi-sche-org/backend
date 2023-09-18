package service

import (
	"context"
	"fmt"
	"time"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/service/internal/converter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type GoogleService interface {
	GetEvents(ctx context.Context, oui entity.OauthUserInfo) ([]entity.CalendarEvent, error)
}

type googleService struct {
	googleCfg *oauth2.Config
}

func NewGoogleService(cfg *config.Config, oar repository.OauthRepository) GoogleService {

	p, err := oar.RegisterProvider(context.Background(), entity.OauthProvider{
		Name:         "google",
		ClientId:     cfg.OAuth.Google.ClientID,
		ClientSecret: cfg.OAuth.Google.ClientSecret,
	})
	if err != nil {
		panic(err)
	}
	gcfg := &oauth2.Config{
		ClientID:     p.ClientId,
		ClientSecret: p.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  fmt.Sprintf("%s/oauth2/google/callback", cfg.BaseURL), // "http://localhost:8080/oauth2/google/callback",
		Scopes:       []string{"openid", "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/calendar.readonly"},
	}
	return &googleService{
		googleCfg: gcfg,
	}
}

func (gs googleService) GetEvents(ctx context.Context, oui entity.OauthUserInfo) ([]entity.CalendarEvent, error) {
	// accessTokenExpires := oui.AccessTokenExpiresAt
	// bufferedExpires := accessTokenExpires.Add(-1 * time.Minute)
	// now := appcontext.Extract(ctx).Now

	// if now.After(bufferedExpires) {
	// 	token := &oauth2.Token{
	// 		AccessToken:  oui.AccessToken,
	// 		RefreshToken: oui.RefreshToken,
	// 		Expiry:       accessTokenExpires,
	// 	}
	// }
	token := converter.OauthUserInfoEntityToOauth2Token(oui)
	tokenSource := gs.googleCfg.TokenSource(ctx, token)
	service, err := v2.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		return []entity.CalendarEvent{}, fmt.Errorf("failed to create google calendar service: %w", err)
	}
	userinfo, err := service.Userinfo.Get().Do()
	if err != nil {
		return []entity.CalendarEvent{}, fmt.Errorf("failed to get userinfo: %w", err)
	}

	calendarService, err := calendar.NewService(
		ctx,
		option.WithTokenSource(tokenSource),
	// option.WithScopes(calendar.CalendarSettingsReadonlyScope),
	)
	if err != nil {
		return []entity.CalendarEvent{}, fmt.Errorf("failed to create google calendar service: %w", err)
	}
	fmt.Printf("calendarService: %+v\n", calendarService)
	fmt.Printf("userinfo: %+v\n", userinfo)

	calendarList, err := calendarService.CalendarList.List().Context(ctx).Do()
	if err != nil {
		return []entity.CalendarEvent{}, fmt.Errorf("failed to get calendar list: %w", err)
	}
	fmt.Printf("calendarList: %+v\n", calendarList)

	timeMin := time.Now().Add(-time.Hour * 24 * 7).Format(time.RFC3339)
	timeMax := time.Now().Add(time.Hour * 24 * 365).Format(time.RFC3339)
	events, err := calendarService.Events.List("primary").SingleEvents(true).OrderBy("startTime").TimeMin(timeMin).TimeMax(timeMax).MaxResults(2500).Context(ctx).Do()
	if err != nil {
		return []entity.CalendarEvent{}, fmt.Errorf("failed to get events: %w", err)
	}
	fmt.Printf("events: %+v\n", events)
	eventsEntity, err := converter.CalendarEventsToEntity(events)
	if err != nil {
		return []entity.CalendarEvent{}, fmt.Errorf("failed to convert events to entity: %w", err)
	}
	fmt.Printf("eventsEntity: %+v\n", eventsEntity)
	return eventsEntity, nil

	// client := gs.googleCfg.Client(ctx, token)
	// panic("unimplemented")

	// client := oc.googleCfg.Client(ctx, token)

	// 	service, err := v2.NewService(ctx, option.WithHTTPClient(client))
	// 	if err != nil {
	// 		return c.JSON(500, "error2")
	// 	}
	// 	// userinfo, err := service.Userinfo.Get().Do()
	// 	userInfo, err := service.Tokeninfo().AccessToken(token.AccessToken).Context(ctx).Do()
	// 	if err != nil {
	// 		return c.JSON(500, "error3")
	// 	}
	// 	// get calendar info
	// 	calendarService, err := calendar.NewService(ctx, option.WithHTTPClient(client), option.WithScopes(calendar.CalendarSettingsReadonlyScope))
	// 	if err != nil {
	// 		return c.JSON(500, "error4")
	// 	}
	// 	calendarList, err := calendarService.CalendarList.List().Context(ctx).Do()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	timeMin := time.Now().Add(-time.Hour * 24 * 7).Format(time.RFC3339)
	// 	timeMax := time.Now().Add(time.Hour * 24 * 365).Format(time.RFC3339)
	// 	events, err := calendarService.Events.List("primary").SingleEvents(true).OrderBy("startTime").TimeMin(timeMin).TimeMax(timeMax).MaxResults(2500).Context(ctx).Do()
	// 	if err != nil {
	// 		return err
	// 		// return c.JSON(500, "error5")
	// 	}
	// 	// events
	// 	return c.Redirect(302, oc.cfg.OAuth.DefaultReturnURL)

	// return nil, nil
}
