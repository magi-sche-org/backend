package router

import (
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/controller"
	"github.com/geekcamp-vol11-team30/backend/logger"
	"github.com/geekcamp-vol11-team30/backend/middleware"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewRouter(
	config *config.Config,
	zlogger *zap.Logger,
	em middleware.ErrorMiddleware,
	atm middleware.AccessTimeMiddleware,
	am middleware.AuthMiddleware,
	uc controller.UserController,
	ac controller.AuthController,
	ec controller.EventController,
) *echo.Echo {
	// TODO: CORSの設定などを足す
	e := echo.New()
	// e.HTTPErrorHandler = em.ErrorHandler
	e.Use(em.ErrorHandler)
	e.Use(atm.Handler)
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	logger.SetRequestLoggerToEcho(e, zlogger)

	e.POST("/users", uc.Register)
	e.POST("/token", ac.CreateUnregisteredUserAndToken)
	e.POST("/token/refresh", ac.RefreshToken)

	eg := e.Group("/events")
	eg.POST("", ec.Create, am.Handler)
	eig := eg.Group("/:event_id")
	eig.GET("", ec.Retrieve)
	// eg.PUT("", ec.Update)
	// eg.DELETE(""", ec.Delete)

	eiag := eig.Group("/user/attend")
	eiag.Use(am.Handler)
	eiag.POST("", ec.Attend)
	eiag.PUT("", ec.Attend)

	// umg := e.Group("/user")
	// ug.GET("", uc.Get)
	// ug.PUT("", uc.Update)
	// ug.DELETE("", uc.Delete)

	// umeg := umg.Group("/events")
	// umeg.GET("", uc.GetEvents)

	return e
}
