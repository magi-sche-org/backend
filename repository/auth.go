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
	FetchRefreshToken(ctx context.Context, token string) (models.RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, token string) error
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

func (ar *authRepository) FetchRefreshToken(ctx context.Context, token string) (models.RefreshToken, error) {
	//get refreshToken
	rt, err := models.RefreshTokens(models.RefreshTokenWhere.Token.EQ(token)).One(ctx, ar.db)
	if err != nil {
		return models.RefreshToken{}, err
	}
	return *rt, nil
}

// DeleteRefreshToken implements AuthRepository.
func (ar *authRepository) DeleteRefreshToken(ctx context.Context, token string) error {
	_, err := models.RefreshTokens(models.RefreshTokenWhere.Token.EQ(token)).DeleteAll(ctx, ar.db)
	if err != nil {
		return err
	}
	return nil
}
