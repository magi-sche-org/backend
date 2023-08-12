package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"go.uber.org/zap"
)

type AuthUsecase interface {
	CreateToken(context.Context, entity.User) (entity.Token, error)
	VerifyAccessToken(ctx context.Context, tokenString string) (userId ulid.ULID, err error)
	RefreshToken(ctx context.Context, refreshToken string) (entity.Token, error)
}

type authUsecase struct {
	cfg       *config.Config
	logger    *zap.Logger
	ar        repository.AuthRepository
	jwtparser *jwt.Parser
}

func NewAuthUsecase(cfg *config.Config, logger *zap.Logger, ar repository.AuthRepository) AuthUsecase {
	parser := newJwtParser()
	return &authUsecase{
		cfg:       cfg,
		logger:    logger,
		ar:        ar,
		jwtparser: parser,
	}
}
func newJwtParser() *jwt.Parser {
	availableMethod := []string{
		jwt.SigningMethodHS256.Name,
	}
	methodOption := jwt.WithValidMethods(availableMethod)
	parser := jwt.NewParser(methodOption)
	return parser
}

// ResponseInChannel implements SlackUsecase.
func (au *authUsecase) CreateToken(ctx context.Context, user entity.User) (entity.Token, error) {
	key := au.cfg.SecretKey
	now := time.Now()
	accessExpire := now.Add(time.Duration(au.cfg.AccessExpireMinutes) * time.Minute)
	userId := user.ID
	sub := util.ULIDToString(userId)

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(accessExpire),
		IssuedAt:  jwt.NewNumericDate(now),
		Subject:   sub,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(key))
	if err != nil {
		return entity.Token{}, err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		return entity.Token{}, err
	}
	refreshExpire := now.Add(time.Duration(au.cfg.RefreshExpireMinutes) * time.Minute)
	err = au.ar.RegisterRefreshToken(ctx, user, refreshToken, refreshExpire)
	if err != nil {
		return entity.Token{}, err
	}
	return entity.Token{
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  accessExpire,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: refreshExpire,
	}, nil

}

func generateRefreshToken() (string, error) {
	// b := make([]byte, 32)
	// _, err := rand.Read(b)
	// if err != nil {
	// 	return "", nil
	// }
	// return hex.EncodeToString(b), nil
	return util.MakeRandomStr(32)
}

// VerifyAccessToken implements AuthUsecase.
func (au *authUsecase) VerifyAccessToken(ctx context.Context, tokenString string) (userId ulid.ULID, err error) {
	token, err := au.jwtparser.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(au.cfg.SecretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			au.logger.Warn("user auth failed", zap.Error(err))
			// return ulid.ULID{}, apperror.NewTokenExpiredError(err, nil)
			return ulid.ULID{}, err
		}
		au.logger.Warn("user auth failed", zap.Error(err))
		return ulid.ULID{}, apperror.NewUnauthorizedError(err, nil, "4000-1")
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	userId, err = util.ULIDFromString(claims.Subject)
	if err != nil {
		au.logger.Error("user auth failed", zap.Error(err))
		return ulid.ULID{}, apperror.NewUnauthorizedError(err, nil, "4000-2")
	}
	return userId, nil
}

// RefreshToken implements AuthUsecase.
func (au *authUsecase) RefreshToken(ctx context.Context, refreshToken string) (entity.Token, error) {
	log.Println(refreshToken)
	refreshTokenRecord, err := au.ar.FetchRefreshToken(ctx, refreshToken)
	log.Println("rrrrrr", refreshTokenRecord)
	if err != nil {
		return entity.Token{}, apperror.NewUnauthorizedError(errors.New("refresh token record not found"), nil, "4000-333")
	}

	// Verify if the refresh token is revoked
	if refreshTokenRecord.Revoked {
		return entity.Token{}, apperror.NewUnauthorizedError(errors.New("refresh token revoked"), nil, "4000-4")
	}

	// Verify if the refresh token has expired
	now := time.Now()
	if refreshTokenRecord.ExpiresAt.Before(now) {
		return entity.Token{}, apperror.NewTokenExpiredError(errors.New("refresh token expired"), nil)
	}

	// Get the user associated with the refresh token
	userID, err := util.ULIDFromString(refreshTokenRecord.UserID)
	if err != nil {
		return entity.Token{}, err
	}

	// Generate new access and refresh tokens
	newAccessToken, err := au.CreateToken(ctx, entity.User{ID: userID})
	if err != nil {
		return entity.Token{}, err
	}

	// Revoke the previous refresh token
	err = au.ar.DeleteRefreshToken(ctx, refreshToken)
	if err != nil {
		return entity.Token{}, err
	}
	return newAccessToken, nil

	// return entity.Token{
	// 	AccessToken:          newAccessToken.AccessToken,
	// 	AccessTokenExpiredAt: time.Now().Add(time.Duration(au.cfg.AccessExpireMinutes) * time.Minute),
	// }, nil
}
