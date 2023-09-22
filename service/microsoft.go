package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	// "errors"
	"fmt"
	"time"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/service/internal/converter"
	"github.com/geekcamp-vol11-team30/backend/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
	// "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	// "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	// auth "github.com/microsoft/kiota-authentication-azure-go"
	// msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	// "github.com/microsoftgraph/msgraph-sdk-go/models"
)

type microsoftService struct {
	// googleCfg *oauth2.Config
	msCfg *oauth2.Config
	oar   repository.OauthRepository
	ur    repository.UserRepository
}

func NewMicrosoftService(cfg *config.Config, oar repository.OauthRepository, ur repository.UserRepository) OauthCalendarService {

	p, err := oar.RegisterProvider(context.Background(), entity.OauthProvider{
		Name:         "microsoft",
		ClientId:     cfg.OAuth.Microsoft.ClientID,
		ClientSecret: cfg.OAuth.Microsoft.ClientSecret,
	})
	if err != nil {
		panic(err)
	}

	msCfg := &oauth2.Config{
		ClientID:     p.ClientId,
		ClientSecret: p.ClientSecret,
		Endpoint:     microsoft.AzureADEndpoint(""),
		RedirectURL:  fmt.Sprintf("%s/oauth2/microsoft/callback", cfg.BaseURL),
		Scopes: []string{
			// "openid",
			"offline_access",
			"Calendars.Read",
			"user.read",
		},
	}
	fmt.Printf("msCfg: %+v\n", msCfg)
	// fmt.Printf("gcfguc: %+v\n", gcfg)
	return &microsoftService{
		// googleCfg: gcfg,
		msCfg: msCfg,
		oar:   oar,
		ur:    ur,
	}
}

// ExchangeToken implements OauthCalendarService.
func (*microsoftService) ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error) {
	panic("unimplemented")
}

// GetAuthURL implements OauthCalendarService.
func (ms *microsoftService) GetAuthURL(ctx context.Context) (url string, state string, err error) {
	state, err = util.MakeRandomStr(32)
	if err != nil {
		return "", "", fmt.Errorf("failed to make random string: %w", err)
	}
	url = ms.msCfg.AuthCodeURL(state,
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)
	return url, state, nil
}

type msUserResponse struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type msCalendarEventResponse struct {
	Subject string `json:"subject"`
}

// GetOrCreateUserByCode implements OauthCalendarService.
func (gs *microsoftService) GetOrCreateUserByCode(ctx context.Context, code string, user *entity.User) (*entity.User, error) {
	token, err := gs.msCfg.Exchange(
		ctx,
		code,
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %w", err)
	}
	tokenSource := gs.msCfg.TokenSource(ctx, token)

	client := gs.msCfg.Client(ctx, token)
	resp, err := client.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user details: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	fmt.Printf("body: %+v\n", string(body))

	var msUser msUserResponse
	if err := json.Unmarshal(body, &msUser); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	fmt.Printf("msUser: %+v\n", msUser)
	// return nil, errors.New("unimplemented")
	uid := msUser.Id

	provider, err := gs.oar.FetchProviderByName(ctx, "microsoft")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch provider by name: %w", err)
	}

	{
		// 既に登録済みのUIDなら，そのユーザーを返しトークン更新（ログイン中は無視）
		oui, err := gs.oar.FetchUserInfoByUid(ctx, provider.ID, uid)
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
	if user != nil { // 既にログイン中
		// IsRegisteredでなければ名前を設定，IsRegisteredへ
		if !user.IsRegistered {
			user.Name = msUser.DisplayName
			user.IsRegistered = true
			err = gs.ur.Update(ctx, *user)
			if err != nil {
				return nil, fmt.Errorf("failed to update user: %w", err)
			}
		}
	} else { // 新規登録
		user = &entity.User{
			Name:         msUser.DisplayName,
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
		ProviderUid:           uid,
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

// // GetPrimaryCalendar implements GoogleService.
// func (gs *googleService) GetPrimaryCalendar(ctx context.Context, oui entity.OauthUserInfo, timeMin *time.Time, timeMax *time.Time) (entity.Calendar, error) {
// 	tokenSource, err := gs.fetchTokenSource(ctx, oui)
// 	if err != nil {
// 		return entity.Calendar{}, fmt.Errorf("failed to fetch token source: %w", err)
// 	}
// 	calendarService, err := gs.fetchCalendarService(ctx, tokenSource)
// 	if err != nil {
// 		return entity.Calendar{}, fmt.Errorf("failed to fetch calendar service: %w", err)
// 	}
// 	events, summary, err := gs.fetchEventsByCalendarID(ctx, calendarService, idPrimary, timeMin, timeMax)
// 	if err != nil {
// 		return entity.Calendar{}, fmt.Errorf("failed to get events: %w", err)
// 	}

// 	calendar, err := converter.GoogleCalendarEventsToEntity(events, idPrimary, summary)
// 	if err != nil {
// 		return entity.Calendar{}, fmt.Errorf("failed to convert events to entity: %w", err)
// 	}
// 	return calendar, nil
// }

// GetPrimaryCalendar implements OauthCalendarService.
func (ms *microsoftService) GetPrimaryCalendar(ctx context.Context, oui entity.OauthUserInfo, timeMin *time.Time, timeMax *time.Time) (entity.Calendar, error) {
	tokenSource, err := ms.fetchTokenSource(ctx, oui)
	if err != nil {
		return entity.Calendar{}, fmt.Errorf("failed to fetch token source: %w", err)
	}
	token, err := tokenSource.Token()
	if err != nil {
		return entity.Calendar{}, fmt.Errorf("failed to get token: %w", err)
	}
	client := ms.msCfg.Client(ctx, token)

	resp, err := client.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		return entity.Calendar{}, fmt.Errorf("failed to fetch calendar: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entity.Calendar{}, fmt.Errorf("received non-OK status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entity.Calendar{}, fmt.Errorf("failed to read response body: %w", err)
	}
	fmt.Printf("body: %+v\n", string(body))

	var msUser msUserResponse
	if err := json.Unmarshal(body, &msUser); err != nil {
		return entity.Calendar{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	panic("unimplemented")

}

func (ms microsoftService) checkAndUpdateRepoToken(ctx context.Context, tokenSource oauth2.TokenSource, oui entity.OauthUserInfo) (entity.OauthUserInfo, error) {
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
	oui, err = ms.oar.UpdateOauthUserInfo(ctx, oui)
	if err != nil {
		return entity.OauthUserInfo{}, fmt.Errorf("failed to update oauth user info: %w", err)
	}
	return oui, nil
}

func (ms microsoftService) fetchTokenSource(ctx context.Context, oui entity.OauthUserInfo) (oauth2.TokenSource, error) {
	token := converter.OauthUserInfoEntityToOauth2Token(oui)
	tokenSource := ms.msCfg.TokenSource(ctx, token)

	_, err := ms.checkAndUpdateRepoToken(ctx, tokenSource, oui)
	if err != nil {
		return nil, fmt.Errorf("failed to check and update repo token: %w", err)
	}
	return tokenSource, nil
}
