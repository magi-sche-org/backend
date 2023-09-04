package usecase

import (
	"context"
	"slices"

	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/validator"
	"github.com/oklog/ulid/v2"
	"golang.org/x/oauth2"
)

type UserUsecase interface {
	CreateAnonymousUser(ctx context.Context) (entity.User, error)
	FindUserByID(ctx context.Context, id ulid.ULID) (entity.User, error)
	Register(ctx context.Context, user entity.User) (entity.UserResponse, error)
	FetchExternalCalendars(ctx context.Context, user entity.User)
}

type userUsecase struct {
	ur  repository.UserRepository
	oar repository.OauthRepository
	uv  validator.UserValidator
}

func NewUserUsecase(ur repository.UserRepository, oar repository.OauthRepository, uv validator.UserValidator) UserUsecase {
	return &userUsecase{
		ur:  ur,
		oar: oar,
		uv:  uv,
	}
}

// CreateAnonymousUser implements UserUsecase.
func (uu *userUsecase) CreateAnonymousUser(ctx context.Context) (entity.User, error) {
	// panic("unimplemented")
	user := entity.User{
		Name:         "",
		IsRegistered: false,
	}
	user, err := uu.ur.Create(ctx, user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// FindUserByID implements UserUsecase.
func (uu *userUsecase) FindUserByID(ctx context.Context, id ulid.ULID) (entity.User, error) {
	user, err := uu.ur.Find(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// Register implements UserUsecase.
func (uu *userUsecase) Register(ctx context.Context, user entity.User) ([]entity.CalendarEvent, error) {
	gp, err := uu.oar.FetchProviderByName(ctx, "google")
	if err != nil {
		return []entity.CalendarEvent{}, err
	}
	// oauis,err := uu.oar.FetchOauthUserInfos(ctx, user)
	// if err != nil {
	// 	return []entity.CalendarEvent{}, err
	// }
	// googlei := slices.IndexFunc(oauis,  func(oaui entity.OauthUserInfo) bool{
	// 	return oaui.
	// })
	oaui, err := uu.oar.FetchOauthUserInfo(ctx)
	token := &oauth2.Token{
		AccessToken: oaui,
	}
	// panic("unimplemented")
	// err := uu.uv.Validate(user)
	// if err != nil {
	// 	return entity.UserResponse{}, err
	// }
	// newUser, err := uu.ur.Create(ctx, user)
	// res := entity.UserResponse{
	// 	// ID:        newUser.ID,
	// 	Name: newUser.Name,
	// 	// SlackID:   newUser.SlackID,
	// 	// CreatedAt: newUser.CreatedAt,
	// 	// UpdatedAt: newUser.UpdatedAt,
	// }
	// return res, err
}
