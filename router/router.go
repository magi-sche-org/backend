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
	oc controller.OauthController,
) *echo.Echo {
	// TODO: CORSの設定などを足す
	e := echo.New()
	// e.HTTPErrorHandler = em.ErrorHandler
	e.Use(em.ErrorHandler)
	e.Use(am.CORSHandler)
	e.Use(am.CSRFHandler)
	e.Use(atm.Handler)

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	e.GET("/csrf", ac.CreateCSRFToken)

	logger.SetRequestLoggerToEcho(e, zlogger)

	e.POST("/users", uc.Register)
	e.POST("/token", ac.CreateUnregisteredUserAndToken)
	e.POST("/token/refresh", ac.RefreshToken)

	og := e.Group("/oauth2")
	og.GET("/google", oc.RedirectToAuthPage)
	og.GET("/google/callback", oc.Callback)

	eg := e.Group("/events")
	eg.Use(am.SessionHandler)
	eg.POST("", ec.Create)
	eig := eg.Group("/:event_id")
	eig.GET("", ec.Retrieve)
	// eg.PUT("", ec.Update)
	// eg.DELETE(""", ec.Delete)

	eiag := eig.Group("/user/answer")
	eiag.Use(am.SessionHandler)
	eiag.GET("", ec.RetrieveUserAnswer)
	eiag.POST("", ec.CreateAnswer)
	eiag.PUT("", ec.CreateAnswer)

	// umg := e.Group("/user")
	// ug.GET("", uc.Get)
	// ug.PUT("", uc.Update)
	// ug.DELETE("", uc.Delete)

	// umeg := umg.Group("/events")
	// umeg.GET("", uc.GetEvents)

	return e
}
