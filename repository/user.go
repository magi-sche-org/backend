package repository

import (
	"context"
	"database/sql"

	"github.com/geekcamp-vol11-team30/backend/db/models"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Find(ctx context.Context, id ulid.ULID) (entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, id string) error
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
	id := util.GenerateULID(ctx)
	u := &models.User{
		ID:   util.ULIDToString(id),
		Name: user.Name,
	}
	err := u.Insert(ctx, ur.db, boil.Infer())
	if err != nil {
		return entity.User{}, err
	}
	return ur.modelToEntity(u)
	// if err != nil {
	// 	return entity.User{}, err
	// }
	// return entity.User{
	// 	ID:           id,
	// 	Name:         u.Name,
	// 	IsRegistered: u.IsRegistered,
	// }, nil
	// panic("unimplemented")
	// now := time.Now()
	// id, err := util.GenerateULID(now)
	// if err != nil {
	// 	return entity.User{}, err
	// }

	// u := &models.User{
	// 	ID:        id.String(),
	// 	Name:      user.Name,
	// 	SlackID:   user.SlackID,
	// 	CreatedAt: now.Format(time.RFC3339Nano),
	// 	UpdatedAt: now.Format(time.RFC3339Nano),
	// }
	// err = u.Insert(ctx, ur.db, boil.Infer())
	// if err != nil {
	// 	return entity.User{}, err
	// }

	// return entity.User{
	// 	ID:        id,
	// 	Name:      u.Name,
	// 	SlackID:   u.SlackID,
	// 	CreatedAt: now,
	// 	UpdatedAt: now,
	// }, nil
}

// Delete implements UserRepository.
func (ur *userRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Find implements UserRepository.
func (ur *userRepository) Find(ctx context.Context, id ulid.ULID) (entity.User, error) {
	m, err := models.FindUser(ctx, ur.db, util.ULIDToString(id))
	if err != nil {
		return entity.User{}, err
	}
	return ur.modelToEntity(m)
}

// FindAll implements UserRepository.
func (ur *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (ur *userRepository) Update(ctx context.Context, user entity.User) error {
	panic("unimplemented")
}

func (ur *userRepository) modelToEntity(m *models.User) (entity.User, error) {
	id, err := util.ULIDFromString(m.ID)
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{
		ID:           id,
		Name:         m.Name,
		IsRegistered: m.IsRegistered,
	}, nil
}
