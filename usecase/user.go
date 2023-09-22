package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/service"
	"github.com/geekcamp-vol11-team30/backend/validator"
	"github.com/oklog/ulid/v2"
)

type UserUsecase interface {
	CreateAnonymousUser(ctx context.Context) (entity.User, error)
	FindUserByID(ctx context.Context, id ulid.ULID) (entity.User, error)

	RetrieveUserProviders(ctx context.Context, user entity.User) ([]entity.OauthProvider, []entity.OauthUserInfo, error)
	// Register(ctx context.Context, user entity.User) ([]entity.CalendarEvent, error)
	FetchExternalCalendars(ctx context.Context, user entity.User, timeMin *time.Time, timeMax *time.Time) ([]entity.Calendar, error)
	ListEvents(ctx context.Context, user entity.User) ([]entity.Event, error)
}

type userUsecase struct {
	ur  repository.UserRepository
	oar repository.OauthRepository
	er  repository.EventRepository
	uv  validator.UserValidator
	gs  service.GoogleService
}

func NewUserUsecase(ur repository.UserRepository, oar repository.OauthRepository, er repository.EventRepository, uv validator.UserValidator, gs service.GoogleService) UserUsecase {
	return &userUsecase{
		ur:  ur,
		oar: oar,
		er:  er,
		uv:  uv,
		gs:  gs,
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

// RetrieveUserProviders implements UserUsecase.
func (uu *userUsecase) RetrieveUserProviders(ctx context.Context, user entity.User) ([]entity.OauthProvider, []entity.OauthUserInfo, error) {
	ouis, err := uu.oar.FetchOauthUserInfos(ctx, user)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch oauth user infos: %w", err)
	}
	ops, err := uu.oar.FetchProviders(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch providers: %w", err)
	}
	return ops, ouis, nil

}

// // Register implements UserUsecase.
// func (uu *userUsecase) Register(ctx context.Context, user entity.User) ([]entity.CalendarEvent, error) {
// 	gp, err := uu.oar.FetchProviderByName(ctx, "google")
// 	if err != nil {
// 		return []entity.CalendarEvent{}, err
// 	}
// 	// oauis,err := uu.oar.FetchOauthUserInfos(ctx, user)
// 	// if err != nil {
// 	// 	return []entity.CalendarEvent{}, err
// 	// }
// 	// googlei := slices.IndexFunc(oauis,  func(oaui entity.OauthUserInfo) bool{
// 	// 	return oaui.
// 	// })
// 	oaui, err := uu.oar.FetchOauthUserInfo(ctx)
// 	token := &oauth2.Token{
// 		AccessToken: oaui,
// 	}
// 	// panic("unimplemented")
// 	// err := uu.uv.Validate(user)
// 	// if err != nil {
// 	// 	return entity.UserResponse{}, err
// 	// }
// 	// newUser, err := uu.ur.Create(ctx, user)
// 	// res := entity.UserResponse{
// 	// 	// ID:        newUser.ID,
// 	// 	Name: newUser.Name,
// 	// 	// SlackID:   newUser.SlackID,
// 	// 	// CreatedAt: newUser.CreatedAt,
// 	// 	// UpdatedAt: newUser.UpdatedAt,
// 	// }
// 	// return res, err
// }

// FetchExternalCalendars implements UserUsecase.
func (uu *userUsecase) FetchExternalCalendars(ctx context.Context, user entity.User, timeMin *time.Time, timeMax *time.Time) ([]entity.Calendar, error) {
	ouis, err := uu.oar.FetchOauthUserInfos(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch oauth user infos: %w", err)
	}
	// fmt.Printf("ouis: %+v\n%+v\n", ouis, ouis[0].Provider)

	// eventsAll := [][]entity.CalendarEvent{}
	calendars := []entity.Calendar{}

	for _, oui := range ouis {
		// fmt.Printf("oui: %+v\n", oui)
		// oui.Provider
		if oui.Provider.Name == "google" {
			calendar, err := uu.gs.GetPrimaryCalendar(ctx, oui, timeMin, timeMax)
			if err != nil {
				return nil, fmt.Errorf("failed to get events: %w", err)
			}
			// fmt.Printf("events: %+v\n", events)
			calendars = append(calendars, calendar)
		}
	}

	return calendars, nil
}

// ListEvents implements UserUsecase.
func (uu *userUsecase) ListEvents(ctx context.Context, user entity.User) ([]entity.Event, error) {
	events, err := uu.er.FetchUserEvents(ctx, nil, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %w", err)
	}
	return events, nil
}
