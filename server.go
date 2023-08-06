package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	e      *echo.Echo
	l      net.Listener
	logger *zap.Logger
}

func NewServer(e *echo.Echo, l net.Listener, logger *zap.Logger) *Server {
	return &Server{
		e:      e,
		l:      l,
		logger: logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	eg, ctx := errgroup.WithContext(ctx)

	s.e.Listener = s.l
	// logger.SetRequestLoggerToEcho(s.e, s.logger)

	s.logger.Info("Server is starting...")
	eg.Go(func() error {
		if err := s.e.Start(""); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	<-ctx.Done()
	s.logger.Info("Server is shutting down...")
	if err := s.e.Shutdown(context.Background()); err != nil {
		s.logger.Error("Server shutdown error", zap.Error(err))
	}
	return eg.Wait()
}
