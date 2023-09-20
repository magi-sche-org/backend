package service

import (
	"context"
	"fmt"
	"time"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/service/internal/converter"
	"github.com/geekcamp-vol11-team30/backend/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

const idPrimary = "primary"
const maxFetchEvents = 5000
const onceFetchEvents = 1000

type GoogleService interface {
	GetGoogleAuthURL(ctx context.Context) (url string, state string, err error)
	ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error)
	GetOrCreateUserByCode(ctx context.Context, code string, user *entity.User) (*entity.User, error)
	GetPrimaryCalendar(ctx context.Context, oui entity.OauthUserInfo, timeMin *time.Time, timeMax *time.Time) (entity.Calendar, error)
}

type googleService struct {
	googleCfg *oauth2.Config
	oar       repository.OauthRepository
	ur        repository.UserRepository
}

func NewGoogleService(cfg *config.Config, oar repository.OauthRepository, ur repository.UserRepository) GoogleService {

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
		Scopes: []string{
			"openid",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/calendar.readonly",
		},
	}
	fmt.Printf("gcfguc: %+v\n", gcfg)
	return &googleService{
		googleCfg: gcfg,
		oar:       oar,
		ur:        ur,
	}
}

// GetGoogleAuthURL implements GoogleService.
func (gs *googleService) GetGoogleAuthURL(ctx context.Context) (url string, state string, err error) {
	state, err = util.MakeRandomStr(32)
	if err != nil {
		return "", "", err
	}
	url = gs.googleCfg.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)
	return url, state, nil
}

// ExchangeToken implements GoogleService.
func (gs *googleService) ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := gs.googleCfg.Exchange(
		ctx,
		code,
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %w", err)
	}
	return token, nil
}

// user infoがあれば，それに紐付いたユーザーを返す。なければ，新規ユーザーを作成して返す
func (gs *googleService) GetOrCreateUserByCode(ctx context.Context, code string, user *entity.User) (*entity.User, error) {
	token, err := gs.googleCfg.Exchange(
		ctx,
		code,
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	tokenSource := gs.googleCfg.TokenSource(ctx, token)

	service, err := v2.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, fmt.Errorf("failed to create google calendar service: %w", err)
	}

	// UIDなど取得
	tokenInfo, err := service.Tokeninfo().AccessToken(token.AccessToken).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get token info: %w", err)
	}

	provider, err := gs.oar.FetchProviderByName(ctx, "google")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch provider by name: %w", err)
	}

	{
		// 既に登録済みのUIDなら，そのユーザーを返しトークン更新（ログイン中は無視）
		oui, err := gs.oar.FetchUserInfoByUid(ctx, provider.ID, tokenInfo.UserId)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch user info by uid: %w", err)
		}

		// 既に登録済みのUIDなら，そのユーザーを返しトークン更新（ログイン中は無視）
		if oui != nil {
			user, err := gs.ur.Find(ctx, oui.UserId)
			if err != nil {
				return nil, fmt.Errorf("failed to find user: %w", err)
			}
			_, err = gs.checkAndUpdateRepoToken(ctx, tokenSource, *oui)
			if err != nil {
				return nil, fmt.Errorf("failed to check and update repo token: %w", err)
			}
			return &user, nil
		}
	}

	// 未登録のUIDなら，新規ユーザー登録
	userInfo, err := service.Userinfo.Get().Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get userinfo: %w", err)
	}

	if user != nil { // 既にログイン中
		// IsRegisteredでなければ名前を設定，IsRegisteredへ
		if !user.IsRegistered {
			user.Name = userInfo.Name
			user.IsRegistered = true
			err = gs.ur.Update(ctx, *user)
			if err != nil {
				return nil, fmt.Errorf("failed to update user: %w", err)
			}
		}
	} else { // 新規登録
		user = &entity.User{
			Name:         userInfo.Name,
			IsRegistered: true,
		}
		_, err = gs.ur.Create(ctx, *user)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	oui := entity.OauthUserInfo{
		UserId:                user.ID,
		ProviderId:            provider.ID,
		ProviderUid:           tokenInfo.UserId,
		AccessToken:           token.AccessToken,
		RefreshToken:          token.RefreshToken,
		AccessTokenExpiresAt:  token.Expiry,
		RefreshTokenExpiresAt: nil,
	}
	_, err = gs.oar.RegisterOauthUserInfo(ctx, oui)
	if err != nil {
		return &entity.User{}, err
	}
	return user, nil

}

// GetPrimaryCalendar implements GoogleService.
func (gs *googleService) GetPrimaryCalendar(ctx context.Context, oui entity.OauthUserInfo, timeMin *time.Time, timeMax *time.Time) (entity.Calendar, error) {
	tokenSource, err := gs.fetchTokenSource(ctx, oui)
	if err != nil {
		return entity.Calendar{}, fmt.Errorf("failed to fetch token source: %w", err)
	}
	calendarService, err := gs.fetchCalendarService(ctx, tokenSource)
	if err != nil {
		return entity.Calendar{}, fmt.Errorf("failed to fetch calendar service: %w", err)
	}
	events, summary, err := gs.fetchEventsByCalendarID(ctx, calendarService, idPrimary, timeMin, timeMax)
	if err != nil {
		return entity.Calendar{}, fmt.Errorf("failed to get events: %w", err)
	}

	calendar, err := converter.GoogleCalendarEventsToEntity(events, idPrimary, summary)
	if err != nil {
		return entity.Calendar{}, fmt.Errorf("failed to convert events to entity: %w", err)
	}
	return calendar, nil
}

// func (gs googleService) GetEvents(ctx context.Context, oui entity.OauthUserInfo, timeMin time.Time, timeMax time.Time) ([]entity.CalendarEvent, error) {
// 	tokenSource, err := gs.fetchTokenSource(ctx, oui)
// 	if err != nil {
// 		return []entity.CalendarEvent{}, fmt.Errorf("failed to fetch token source: %w", err)
// 	}

// 	calendarService, err := gs.fetchCalendarService(ctx, tokenSource)
// 	if err != nil {
// 		return []entity.CalendarEvent{}, fmt.Errorf("failed to fetch calendar service: %w", err)
// 	}

// 	calendarList, err := gs.fetchCalendarList(ctx, calendarService)
// 	if err != nil {
// 		return []entity.CalendarEvent{}, fmt.Errorf("failed to get calendar list: %w", err)
// 	}
// 	fmt.Printf("calendarList: %+v\n", calendarList)

// 	events, err := gs.fetchEventsByCalendarID(ctx, calendarService, idPrimary, timeMin, timeMax)
// 	if err != nil {
// 		return []entity.CalendarEvent{}, fmt.Errorf("failed to get events: %w", err)
// 	}

// 	eventsEntity, err := converter.GoogleCalendarEventsToEntity(events, idPrimary)
// 	if err != nil {
// 		return []entity.CalendarEvent{}, fmt.Errorf("failed to convert events to entity: %w", err)
// 	}

// 	return eventsEntity, nil
// }

func (gs googleService) fetchTokenSource(ctx context.Context, oui entity.OauthUserInfo) (oauth2.TokenSource, error) {
	token := converter.OauthUserInfoEntityToOauth2Token(oui)
	tokenSource := gs.googleCfg.TokenSource(ctx, token)

	_, err := gs.checkAndUpdateRepoToken(ctx, tokenSource, oui)
	if err != nil {
		return nil, fmt.Errorf("failed to check and update repo token: %w", err)
	}
	return tokenSource, nil
}

func (gs googleService) checkAndUpdateRepoToken(ctx context.Context, tokenSource oauth2.TokenSource, oui entity.OauthUserInfo) (entity.OauthUserInfo, error) {
	token, err := tokenSource.Token()
	if err != nil {
		return oui, fmt.Errorf("failed to get new token: %w", err)
	}

	if oui.AccessToken == token.AccessToken {
		return oui, nil
	}
	oui.AccessToken = token.AccessToken
	oui.AccessTokenExpiresAt = token.Expiry
	oui.RefreshToken = token.RefreshToken
	oui, err = gs.oar.UpdateOauthUserInfo(ctx, oui)
	if err != nil {
		return entity.OauthUserInfo{}, fmt.Errorf("failed to update oauth user info: %w", err)
	}
	return oui, nil
}

// func (gs googleService) fetchUserInfo(ctx context.Context, tokenSource oauth2.TokenSource) (*v2.Userinfo, error) {
// 	service, err := v2.NewService(ctx, option.WithTokenSource(tokenSource))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create google calendar service: %w", err)
// 	}
// 	userInfo, err := service.Userinfo.Get().Do()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get userinfo: %w", err)
// 	}
// 	return userInfo, nil
// }

func (gs googleService) fetchCalendarService(ctx context.Context, tokenSource oauth2.TokenSource) (*calendar.Service, error) {
	calendarService, err := calendar.NewService(
		ctx,
		option.WithTokenSource(tokenSource),
	// option.WithScopes(calendar.CalendarSettingsReadonlyScope),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create google calendar service: %w", err)
	}
	return calendarService, nil
}

// func (gs googleService) fetchCalendarList(ctx context.Context, service *calendar.Service) (*calendar.CalendarList, error) {
// 	calendarList, err := service.CalendarList.List().Context(ctx).Do()

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get calendar list: %w", err)
// 	}
// 	return calendarList, nil
// }

func (gs googleService) fetchEventsByCalendarID(ctx context.Context, service *calendar.Service, id string, timeMin *time.Time, timeMax *time.Time) ([]*calendar.Event, string, error) {
	if timeMin == nil {
		t := time.Now().Add(-time.Hour * 24 * 7)
		timeMin = &t
	}
	if timeMax == nil {
		t := time.Now().Add(time.Hour * 24 * 100)
		timeMax = &t
	}
	summary := ""
	nextPageToken := ""
	items := []*calendar.Event{}
	for {
		fmt.Printf("nextPageToken: %+v\n", nextPageToken)
		events, err := service.Events.List(id).
			EventTypes("default", "focusTime", "outOfOffice").
			SingleEvents(true).
			OrderBy("startTime").
			TimeMin(timeMin.Format(time.RFC3339)).
			TimeMax(timeMax.Format(time.RFC3339)).
			PageToken(nextPageToken).
			MaxResults(onceFetchEvents).
			Context(ctx).Do()
		if err != nil {
			return nil, "", fmt.Errorf("failed to get events: %w", err)
		}

		items = append(items, events.Items...)

		summary = events.Summary
		nextPageToken = events.NextPageToken
		if nextPageToken == "" || len(items) >= maxFetchEvents {
			break
		}
	}
	return items, summary, nil
}
