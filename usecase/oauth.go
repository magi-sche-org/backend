package usecase

import (
	"context"
	"fmt"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/service"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OauthUsecase interface {
	GetGoogleAuthURL(ctx context.Context) (string, string, error)
	LoginGoogleWithCode(ctx context.Context, user *entity.User, code string) (entity.User, error)
}

type oauthUsecase struct {
	cfg       *config.Config
	googleCfg *oauth2.Config
	oar       repository.OauthRepository
	ur        repository.UserRepository
	uu        UserUsecase
	gs        service.GoogleService
}

func NewOauthUsecase(cfg *config.Config, oar repository.OauthRepository, ur repository.UserRepository, gs service.GoogleService, uu UserUsecase) OauthUsecase {
	gcfg := &oauth2.Config{
		ClientID:     cfg.OAuth.Google.ClientID,
		ClientSecret: cfg.OAuth.Google.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  fmt.Sprintf("%s/oauth2/google/callback", cfg.BaseURL), // "http://localhost:8080/oauth2/google/callback",
		Scopes: []string{
			"openid",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/calendar.readonly",
		},
	}
	// fmt.Printf("gcfguc: %+v\n", gcfg)
	return &oauthUsecase{
		cfg:       cfg,
		googleCfg: gcfg,
		oar:       oar,
		ur:        ur,
		uu:        uu,
		gs:        gs,
	}
}

// GetGoogleAuthURL implements OauthUsecase.
func (oau *oauthUsecase) GetGoogleAuthURL(ctx context.Context) (url string, state string, err error) {
	url, state, err = oau.gs.GetGoogleAuthURL(ctx)
	if err != nil {
		return "", "", fmt.Errorf("failed to get google auth url: %w", err)
	}
	return url, state, nil
}

// LoginGoogleWithCode implements OauthUsecase.
func (oau *oauthUsecase) LoginGoogleWithCode(ctx context.Context, user *entity.User, code string) (entity.User, error) {
	user, err := oau.gs.GetOrCreateUserByCode(ctx, code, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to login google with code: %w", err)
	}
	return *user, nil
}
