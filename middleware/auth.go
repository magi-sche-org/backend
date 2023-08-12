//go:generate mockgen -source=./auth.go -destination=./mock/auth.go -package=mockmiddleware
package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/geekcamp-vol11-team30/backend/appcontext"
	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oklog/ulid/v2"
	"go.uber.org/zap"
)

type AuthMiddleware interface {
	CORSHandler(next echo.HandlerFunc) echo.HandlerFunc
	CSRFHandler(next echo.HandlerFunc) echo.HandlerFunc
	SessionHandler(next echo.HandlerFunc) echo.HandlerFunc
	IfLoginSessionHandler(next echo.HandlerFunc) echo.HandlerFunc
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
	return middleware.CORSWithConfig(am.corsCfg)(next)
}

// CSRFHijackHandler implements AuthMiddleware.
func (am *authMiddleware) CSRFHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.CSRFWithConfig(am.csrfCfg)(next)
}

// SessionHandler implements AuthMiddleware.
func (am *authMiddleware) SessionHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		tokenCookie, err := c.Cookie("accessToken")

		needCheckRefreshToken := false
		userId := ulid.ULID{}

		if err == nil { // access tokenがあるとき
			tokenString := tokenCookie.Value
			if tokenString == "" { // 空ならぶっとばす
				am.logger.Warn("user auth failed", zap.Error(errors.New("token is empty")))
				return apperror.NewUnauthorizedError(errors.New("token is empty"), nil, "4000-02")
			}
			log.Println(tokenString)
			userId, err = am.verifyAccessToken(ctx, tokenString)
			if err != nil {
				if !errors.Is(err, jwt.ErrTokenExpired) {
					return err
				}
				needCheckRefreshToken = true
			}
		} else { // access token がない・あるいは読み込みエラーのとき
			if !errors.Is(err, http.ErrNoCookie) { // cookieが無い以外のエラー
				return apperror.NewInternalError(err, nil, "access token cookie fetch error")
			}
			am.logger.Info("access token not found in cookie. try to check refresh token")
			needCheckRefreshToken = true

		}
		if needCheckRefreshToken {
			refreshCookie, err := c.Cookie("refreshToken")
			if err != nil {
				am.logger.Warn("user auth failed", zap.Error(errors.New("refresh token must be set")))
				return apperror.NewUnauthorizedError(errors.New("refresh token must be set"), nil, "4000-03")
			}
			refreshTokenString := refreshCookie.Value
			token, err := am.au.RefreshToken(ctx, refreshTokenString)
			if err != nil {
				return err
			}
			util.SetTokenCookie(c, *am.cfg, token)
			userId, err = am.verifyAccessToken(ctx, token.AccessToken)
			if err != nil {
				return err
			}
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

// IfLoginSessionHandler implements AuthMiddleware.
func (am *authMiddleware) IfLoginSessionHandler(next echo.HandlerFunc) echo.HandlerFunc {
	// atc, err := c.Cookie("accessToken")
	return func(c echo.Context) error {
		_, err := c.Cookie("accessToken")
		if err != nil {
			if !errors.Is(err, http.ErrNoCookie) {
				// return next(c)
				return err
			}
			// return err
		}
		_, err = c.Cookie("refreshToken")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				return next(c)
				// return err
			}
			return err
		}
		return am.SessionHandler(next)(c)
	}
}

func (am *authMiddleware) verifyAccessToken(ctx context.Context, tokenString string) (userId ulid.ULID, err error) {
	userId, err = am.au.VerifyAccessToken(ctx, tokenString) // check
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			err = fmt.Errorf("user auth failed, invalid access token: %w", err)
			return ulid.ULID{}, apperror.NewUnauthorizedError(err, nil, "4000-03")
		} else { // その他はエラー
			return ulid.ULID{}, err
		}
	}
	return userId, nil
}
