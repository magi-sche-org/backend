package usecase

import (
	"context"

	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/validator"
	"github.com/oklog/ulid/v2"
)

type UserUsecase interface {
	CreateAnonymousUser(ctx context.Context) (entity.User, error)
	FindUserByID(ctx context.Context, id ulid.ULID) (entity.User, error)
	Register(ctx context.Context, user entity.User) (entity.UserResponse, error)
}

type userUsecase struct {
	ur repository.UserRepository
	uv validator.UserValidator
}

func NewUserUsecase(ur repository.UserRepository, uv validator.UserValidator) UserUsecase {
	return &userUsecase{
		ur: ur,
		uv: uv,
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
func (uu *userUsecase) Register(ctx context.Context, user entity.User) (entity.UserResponse, error) {
	panic("unimplemented")
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
