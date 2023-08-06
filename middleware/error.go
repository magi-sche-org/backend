//go:generate mockgen -source=./error.go -destination=./mock/error.go -package=mockmiddleware
package middleware

import (
	"net/http"

	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ErrorMiddleware interface {
	// ErrorHandler(err error, c echo.Context)
	ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc
}

type errorMiddleware struct {
	logger *zap.Logger
}

func NewErrorMiddleware(logger *zap.Logger, uu usecase.UserUsecase) ErrorMiddleware {
	return &errorMiddleware{
		logger: logger,
	}
}

// Handler implements ErrorMiddleware.
// func (*errorMiddleware) ErrorHandler(err error, c echo.Context) {
// 	if c.Response().Committed {
// 		return
// 	}
// 	log.Println("errorMiddleware")
// 	// return func(c echo.Context) error {
// 	// err := next(c)
// 	if err != nil {
// 		if he, ok := err.(*echo.HTTPError); ok {
// 			c.JSON(he.Code, apperror.NewUnknownError(err, nil))
// 		}
// 		if ae, ok := err.(*apperror.AppError); ok {
// 			c.JSON(ae.StatusCode, ae)
// 		}
// 		// return err
// 	}
// 	// echo.DefaultHTTPErrorHandler(err, c)
// 	// return nil
// 	// }
// }

func (em *errorMiddleware) ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			em.logger.Warn("errorMiddleware", zap.Error(err))
			if he, ok := err.(*echo.HTTPError); ok {
				if he.Code == http.StatusNotFound {
					return c.JSON(he.Code, apperror.NewNotFoundError(err, nil))
				}
				return c.JSON(he.Code, apperror.NewEchoHttpError(he.Code, he.Message, he.Internal))

			}
			if ae, ok := err.(*apperror.AppError); ok {
				return c.JSON(ae.StatusCode, ae)
			}
			return c.JSON(http.StatusInternalServerError, apperror.NewUnknownError(err, nil))
		}
		return nil
	}
}
