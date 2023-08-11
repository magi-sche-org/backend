//go:generate mockgen -source=./auth.go -destination=./mock/auth.go -package=mockmiddleware
package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/geekcamp-vol11-team30/backend/appcontext"
	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type AuthMiddleware interface {
	CORSHandler(next echo.HandlerFunc) echo.HandlerFunc
	CSRFHandler(next echo.HandlerFunc) echo.HandlerFunc
	SessionHandler(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware struct {
	cfg    *config.Config
	logger *zap.Logger
	au     usecase.AuthUsecase
	uu     usecase.UserUsecase

	corsCfg middleware.CORSConfig
	csrfCfg middleware.CSRFConfig
}

func NewAuthMiddleware(cfg *config.Config, logger *zap.Logger, au usecase.AuthUsecase, uu usecase.UserUsecase) AuthMiddleware {
	corsCfg := middleware.CORSConfig{
		Skipper: func(c echo.Context) bool {
			return cfg.CSRF.Disabled
		},
		AllowOrigins: cfg.CORS.Origins,
		// AllowOriginFunc: func(origin string) (bool, error) {
		// },
		AllowMethods: []string{
			echo.GET, echo.PUT, echo.POST, echo.DELETE,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken,
		},
		AllowCredentials: true,
		// UnsafeWildcardOriginWithAllowCredentials: false,
		// ExposeHeaders:                            []string{},
		// MaxAge:                                   0,
	}
	csrfCfg := middleware.CSRFConfig{
		Skipper: func(c echo.Context) bool {
			// if c.Request().Method == http.MethodPost && c.Path() == "/token" {
			// 	return true
			// }
			return cfg.CSRF.Disabled
		},
		// TokenLength:    0,
		// TokenLookup:    "",
		ContextKey:     "csrf",
		CookieName:     "csrf",
		CookieDomain:   cfg.CSRF.Domain,
		CookiePath:     "/",
		CookieMaxAge:   int(time.Duration(time.Duration(cfg.RefreshExpireMinutes) * time.Minute).Seconds()),
		CookieSecure:   cfg.Env != "dev",
		CookieHTTPOnly: cfg.CSRF.HttpOnly,
		CookieSameSite: http.SameSite(cfg.CSRF.SameSite),
		// ErrorHandler: ,
		ErrorHandler: func(err error, c echo.Context) error {
			return apperror.NewMissingCSRFTokenError(err)
		},
	}

	return &authMiddleware{
		cfg:     cfg,
		logger:  logger,
		au:      au,
		uu:      uu,
		corsCfg: corsCfg,
		csrfCfg: csrfCfg,
	}
}

// CORSHandler implements AuthMiddleware.
func (am *authMiddleware) CORSHandler(next echo.HandlerFunc) echo.HandlerFunc {
	// panic("unimplemented")
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	})(next)
}

// CSRFHijackHandler implements AuthMiddleware.
func (am *authMiddleware) CSRFHandler(next echo.HandlerFunc) echo.HandlerFunc {
	// panic("unimplemented")
	// err := next()
	// return err
	return middleware.CSRFWithConfig(am.csrfCfg)(next)
}

// SessionHandler implements AuthMiddleware.
func (am *authMiddleware) SessionHandler(next echo.HandlerFunc) echo.HandlerFunc {
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
