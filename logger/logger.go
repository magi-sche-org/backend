package logger

import (
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func NewLogger(config config.Config) (*zap.Logger, error) {
	if config.Env == "dev" {
		return zap.NewDevelopment()
	} else {
		return zap.NewProduction()
	}
}

// // 過剰なまでのmiddlewareの設定
// func SetRequestLoggerToEcho(e *echo.Echo, logger *zap.Logger) {
// 	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
// 		LogLatency:       true,
// 		LogProtocol:      true,
// 		LogRemoteIP:      true,
// 		LogHost:          true,
// 		LogMethod:        true,
// 		LogURI:           true,
// 		LogURIPath:       true,
// 		LogRoutePath:     true,
// 		LogRequestID:     true,
// 		LogReferer:       true,
// 		LogUserAgent:     true,
// 		LogStatus:        true,
// 		LogError:         true,
// 		LogContentLength: true,
// 		LogResponseSize:  true,
// 		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
// 			logger.Info("request",
// 				zap.String("protocol", v.Protocol),
// 				zap.String("remote_ip", v.RemoteIP),
// 				zap.String("host", v.Host),
// 				zap.String("method", v.Method),
// 				zap.String("uri", v.URI),
// 				zap.String("user_agent", v.UserAgent),
// 				zap.Int("status", v.Status),
// 				zap.Duration("latency", v.Latency),
// 				zap.String("latency_human", v.Latency.String()),
// 				zap.String("content_length", v.ContentLength),
// 				zap.String("referer", v.Referer),
// 				zap.String("request_id", v.RequestID),
// 				zap.String("route_path", v.RoutePath),
// 				zap.String("uri_path", v.URIPath),
// 				zap.Int64("response_size", v.ResponseSize),
// 				zap.Error(v.Error),
// 			)
// 			return nil
// 		},
// 	}))
// }

// 過剰なまでのmiddlewareの設定
func SetRequestLoggerToEcho(e *echo.Echo, logger *zap.Logger) {
	cfg := middleware.RequestLoggerConfig{
		LogLatency:       true,
		LogProtocol:      true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogURIPath:       true,
		LogRoutePath:     true,
		LogRequestID:     true,
		LogReferer:       true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogError:         true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("protocol", v.Protocol),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("host", v.Host),
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.String("user_agent", v.UserAgent),
				zap.Int("status", v.Status),
				zap.Duration("latency", v.Latency),
				zap.String("latency_human", v.Latency.String()),
				zap.String("content_length", v.ContentLength),
				zap.String("referer", v.Referer),
				zap.String("request_id", v.RequestID),
				zap.String("route_path", v.RoutePath),
				zap.String("uri_path", v.URIPath),
				zap.Int64("response_size", v.ResponseSize),
				zap.Error(v.Error),
			)
			return nil
		},
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// if c.Path() == "/health" {
			// 	return next(c)
			// }
			return middleware.RequestLoggerWithConfig(cfg)(next)(c)
		}
	})
}

// 過剰なまでのmiddlewareの設定
