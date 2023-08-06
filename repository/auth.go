package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/geekcamp-vol11-team30/backend/db/models"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AuthRepository interface {
	// tokenを登録するお
	RegisterRefreshToken(ctx context.Context, user entity.User, token string, expiresAt time.Time) error
	UpdateRefreshToken(ctx context.Context, user entity.User, token string, expiresAt time.Time) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

// RegisterRefreshToken implements AuthRepository.
func (ar *authRepository) RegisterRefreshToken(ctx context.Context, user entity.User, token string, expiresAt time.Time) error {
	id := util.GenerateULID(ctx)
	rt := &models.RefreshToken{
		ID:        util.ULIDToString(id),
		UserID:    util.ULIDToString(user.ID),
		Token:     token,
		ExpiresAt: expiresAt,
		Revoked:   false,
	}
	err := rt.Insert(ctx, ar.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

// UpdateRefreshToken implements AuthRepository.
func (*authRepository) UpdateRefreshToken(ctx context.Context, user entity.User, token string, expiresAt time.Time) error {
	// TODO: implement
	panic("unimplemented")
}
