package usecase

import (
	"context"
	"fmt"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/service"
)

type OauthUsecase interface {
	GetAuthURL(ctx context.Context, provider string) (string, string, error)
	LoginWithCode(ctx context.Context, provider string, user *entity.User, code string) (entity.User, error)
}

type oauthUsecase struct {
	cfg *config.Config
	oar repository.OauthRepository
	ur  repository.UserRepository
	uu  UserUsecase
	gs  service.OauthCalendarService
	ms  service.OauthCalendarService
}

func NewOauthUsecase(cfg *config.Config, oar repository.OauthRepository, ur repository.UserRepository, gs service.OauthCalendarService, ms service.OauthCalendarService, uu UserUsecase) OauthUsecase {
	return &oauthUsecase{
		cfg: cfg,
		oar: oar,
		ur:  ur,
		uu:  uu,
		gs:  gs,
		ms:  ms,
	}
}

// GetAuthURL implements OauthUsecase.
func (oau *oauthUsecase) GetAuthURL(ctx context.Context, provider string) (url string, state string, err error) {
	if provider == "google" {
		url, state, err = oau.gs.GetAuthURL(ctx)

		if err != nil {
			return "", "", fmt.Errorf("failed to get google auth url: %w", err)
		}
		return url, state, nil
	} else if provider == "microsoft" {
		url, state, err = oau.ms.GetAuthURL(ctx)
		if err != nil {
			return "", "", fmt.Errorf("failed to get microsoft auth url: %w", err)
		}
		return url, state, nil
	}

	return "", "", fmt.Errorf("provider not found")

}

// LoginWithCode implements OauthUsecase.
func (oau *oauthUsecase) LoginWithCode(ctx context.Context, provider string, user *entity.User, code string) (entity.User, error) {
	if provider == "google" {
		user, err := oau.gs.GetOrCreateUserByCode(ctx, code, user)
		if err != nil {
			return entity.User{}, fmt.Errorf("failed to login google with code: %w", err)
		}
		return *user, nil
	} else if provider == "microsoft" {
		user, err := oau.ms.GetOrCreateUserByCode(ctx, code, user)
		if err != nil {
			return entity.User{}, fmt.Errorf("failed to login microsoft with code: %w", err)
		}
		return *user, nil
	}
	return entity.User{}, fmt.Errorf("provider not found")

}
