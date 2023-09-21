package router

import (
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/controller"
	"github.com/geekcamp-vol11-team30/backend/logger"
	"github.com/geekcamp-vol11-team30/backend/middleware"
	"github.com/geekcamp-vol11-team30/backend/util"
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
	e := echo.New()
	// e.HTTPErrorHandler = em.ErrorHandler
	e.Use(em.ErrorHandler)
	e.Use(am.CORSHandler)
	e.Use(am.CSRFHandler)
	e.Use(atm.Handler)

	e.GET("/health", func(c echo.Context) error {
		return util.JSONResponse(c, 200, "OK")
	})
	e.GET("/csrf", ac.CreateCSRFToken)

	logger.SetRequestLoggerToEcho(e, zlogger)

	// ログイン中ユーザー関連
	eug := e.Group("/user")
	eug.Use(am.SessionHandler)
	eug.GET("", uc.Get)
	// eug.GET("/events", uc.GetEvents)
	eug.GET("/external/calendars", uc.GetExternalCalendars)

	// // ユーザー登録
	// e.POST("/users", uc.Register)
	// 匿名ユーザー登録
	e.POST("/token", ac.CreateUnregisteredUserAndToken)
	e.POST("/token/refresh", ac.RefreshToken)
	e.POST("/logout", ac.Logout)

	// Oauth関連
	og := e.Group("/oauth2")
	// og.Use(am.IfLoginSessionHandler)
	og.GET("/google", oc.RedirectToAuthPage, am.IfLoginSessionHandler)
	og.GET("/google/callback", oc.Callback, am.SessionHandler)

	// イベント関連
	eg := e.Group("/events")
	eg.Use(am.SessionHandler)
	// イベント作成
	eg.POST("", ec.Create)

	eig := eg.Group("/:event_id")
	// 特定イベント情報取得
	eig.GET("", ec.Retrieve)
	// eg.PUT("", ec.Update)
	// eg.DELETE(""", ec.Delete)

	eiag := eig.Group("/user/answer")
	// eiag.Use(am.SessionHandler)
	// 特定イベントのユーザーの回答取得
	eiag.GET("", ec.RetrieveUserAnswer)
	// 特定イベントのユーザーの回答作成
	eiag.POST("", ec.CreateAnswer)
	// 特定イベントのユーザーの回答更新（POSTと同じ）
	eiag.PUT("", ec.CreateAnswer)

	// umg := e.Group("/user")
	// ug.GET("", uc.Get)
	// ug.PUT("", uc.Update)
	// ug.DELETE("", uc.Delete)

	// umeg := umg.Group("/events")
	// umeg.GET("", uc.GetEvents)

	return e
}
