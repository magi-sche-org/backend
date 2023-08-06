//go:generate mockgen -source=./auth.go -destination=./mock/auth.go -package=mockmiddleware
package middleware

import (
	"errors"
	"fmt"
	"log"

	"github.com/geekcamp-vol11-team30/backend/appcontext"
	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AuthMiddleware interface {
	Handler(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware struct {
	cfg    *config.Config
	logger *zap.Logger
	au     usecase.AuthUsecase
	uu     usecase.UserUsecase
}

func NewAuthMiddleware(cfg *config.Config, logger *zap.Logger, au usecase.AuthUsecase, uu usecase.UserUsecase) AuthMiddleware {
	return &authMiddleware{
		cfg:    cfg,
		logger: logger,
		au:     au,
		uu:     uu,
	}
}

// Handler implements AuthMiddleware.
func (am *authMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		// Bearer token format: Bearer <token>
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			am.logger.Warn("user auth failed", zap.Error(errors.New("token is empty")))
			return apperror.NewUnauthorizedError(errors.New("token is empty"), nil, "4000-01")
		}
		if len(tokenString) < len("Bearer ") {
			am.logger.Warn("user auth failed", zap.Error(errors.New("token is invalid")))
			return apperror.NewUnauthorizedError(errors.New("token is invalid"), nil, "4000-02")
		}
		if tokenString[:len("Bearer ")] != "Bearer " {
			am.logger.Warn("user auth failed", zap.Error(errors.New("token is invalid")))
			return apperror.NewUnauthorizedError(errors.New("token is invalid"), nil, "4000-03")
		}
		tokenString = tokenString[len("Bearer "):]
		log.Println(tokenString)
		userId, err := am.au.VerifyAccessToken(ctx, tokenString)
		if err != nil {
			return err
		}

		user, err := am.uu.FindUserByID(ctx, userId)
		if err != nil {
			if aerr, ok := err.(*apperror.AppError); ok && aerr.StatusCode == 404 {
				return apperror.NewUnauthorizedError(fmt.Errorf("user not found: %w", err), nil, "4000-04")
			}
			return err
		}

		actx := appcontext.Extract(ctx)
		actx.SetUser(user)
		appcontext.Set(ctx, actx)
		log.Println(user)

		err = next(c)
		return err
	}
}
