package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type OauthUsecase interface {
	GetGoogleAuthURL(ctx context.Context) (string, string, error)
	LoginGoogleWithCode(ctx context.Context, code string) (*oauth2.Token, error)
	FetchAndRegisterOauthUserInfo(ctx context.Context, token *oauth2.Token, user *entity.User) (entity.User, error)
}

type oauthUsecase struct {
	cfg       *config.Config
	googleCfg *oauth2.Config
	oar       repository.OauthRepository
	ur        repository.UserRepository
	uu        UserUsecase
}

func NewOauthUsecase(cfg *config.Config, oar repository.OauthRepository, ur repository.UserRepository, uu UserUsecase) OauthUsecase {
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
	fmt.Printf("gcfguc: %+v\n", gcfg)
	return &oauthUsecase{
		cfg:       cfg,
		googleCfg: gcfg,
		oar:       oar,
		ur:        ur,
		uu:        uu,
	}
}

// GetGoogleAuthURL implements OauthUsecase.
func (oau *oauthUsecase) GetGoogleAuthURL(ctx context.Context) (url string, state string, err error) {
	state, err = util.MakeRandomStr(32)
	if err != nil {
		return "", "", err
	}
	url = oau.googleCfg.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)
	return url, state, nil
}

// LoginGoogleWithCode implements OauthUsecase.
func (oau *oauthUsecase) LoginGoogleWithCode(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := oau.googleCfg.Exchange(
		ctx,
		code,
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to login google with code(%s): %w", code, err)
	}

	return token, nil
}

// user infoがあれば，それに紐付いたユーザーを返す。なければ，新規ユーザーを作成して返す
func (oau *oauthUsecase) FetchAndRegisterOauthUserInfo(ctx context.Context, token *oauth2.Token, targetUser *entity.User) (entity.User, error) {
	client := oau.googleCfg.Client(ctx, token)
	service, err := v2.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return entity.User{}, err
	}
	userInfo, err := service.Tokeninfo().AccessToken(token.AccessToken).Context(ctx).Do()
	if err != nil {
		return entity.User{}, err
	}

	{
		op, err := oau.oar.FetchProviderByName(ctx, "google")
		if err != nil {
			return entity.User{}, err
		}
		oaui, err := oau.oar.FetchUserInfoByUid(ctx, op.ID, userInfo.UserId)
		if err != nil {
			return entity.User{}, err
		}
		if oaui != nil {
			user, err := oau.ur.Find(ctx, oaui.UserId)
			return user, err
			// return entity.User{}, fmt.Errorf("already registered")
		}
	}
	fmt.Printf("userInfo: %#v\n", userInfo)

	user := entity.User{
		Name:         userInfo.Email,
		IsRegistered: true,
	}
	if targetUser == nil {
		log.Println(user, targetUser, "xxxxxxxxxxxxxxxxxxxxxxxxxxxx")
		// user, err = oau.uu.CreateAnonymousUser(ctx)
		user, err = oau.ur.Create(ctx, user)
		if err != nil {
			return entity.User{}, err
		}
	} else {
		log.Println(user, targetUser, "hhhhhhhhhhhhhhhhhhhhhhhhh")
		user = *targetUser
		user.IsRegistered = true
		err := oau.ur.Update(ctx, user)
		if err != nil {
			return entity.User{}, err
		}
	}
	provider, err := oau.oar.FetchProviderByName(ctx, "google")
	if err != nil {
		return entity.User{}, err
	}
	oaui := entity.OauthUserInfo{
		UserId:                user.ID,
		ProviderId:            provider.ID,
		ProviderUid:           userInfo.UserId,
		AccessToken:           token.AccessToken,
		RefreshToken:          token.RefreshToken,
		AccessTokenExpiresAt:  token.Expiry,
		RefreshTokenExpiresAt: nil,
	}
	_, err = oau.oar.RegisterOauthUserInfo(ctx, oaui)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
