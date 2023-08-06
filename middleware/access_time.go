//go:generate mockgen -source=./access_time.go -destination=./mock/access_time.go -package=mockmiddleware
package middleware

import (
	"time"

	"github.com/geekcamp-vol11-team30/backend/appcontext"
	"github.com/labstack/echo/v4"
)

type AccessTimeMiddleware interface {
	Handler(next echo.HandlerFunc) echo.HandlerFunc
}

type accessTimeMiddleware struct {
	// logger *zap.Logger
}

func NewAccessTimeMiddleware() AccessTimeMiddleware {
	return &accessTimeMiddleware{}
}

// Handler implements AuthMiddleware.
func (am *accessTimeMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		now := time.Now()
		ctx := c.Request().Context()
		ctx = appcontext.Set(ctx, &appcontext.AppContext{
			Now: now,
		})
		c.SetRequest(c.Request().WithContext(ctx))
		err := next(c)
		return err
	}
}
