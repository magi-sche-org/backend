package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository/internal/converter"
	"github.com/geekcamp-vol11-team30/backend/repository/internal/models"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Find(ctx context.Context, id ulid.ULID) (entity.User, error)
	// FindAll(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, user entity.User) error
	// Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create implements UserRepository.
func (ur *userRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	user.ID = util.GenerateULID(ctx)
	um := converter.UserEntityToModel(ctx, user)

	err := um.Insert(ctx, ur.db, boil.Infer())
	if err != nil {
		return entity.User{}, err
	}

	u, err := converter.UserModelToEntity(ctx, um)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to convert UserModel to entity on Create: %w", err)
	}
	return u, nil
}

// // Delete implements UserRepository.
// func (ur *userRepository) Delete(ctx context.Context, id string) error {
// 	panic("unimplemented")
// }

// Find implements UserRepository.
func (ur *userRepository) Find(ctx context.Context, id ulid.ULID) (entity.User, error) {
	um, err := models.FindUser(ctx, ur.db, util.ULIDToString(id))
	if err != nil {
		return entity.User{}, err
	}

	u, err := converter.UserModelToEntity(ctx, um)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to convert UserModel to entity on Find: %w", err)
	}
	return u, nil
}

// // FindAll implements UserRepository.
// func (ur *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
// 	panic("unimplemented")
// }

// Update implements UserRepository.
func (ur *userRepository) Update(ctx context.Context, user entity.User) error {
	um := converter.UserEntityToModel(ctx, user)

	err := um.Upsert(ctx, ur.db, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
