package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func TestServer_Run(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)

	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	eg.Go(func() error {
		s := NewServer(e, l, zap.NewExample())
		return s.Run(ctx)
	})

	url := fmt.Sprintf("http://%s/health", l.Addr().String())
	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if rsp.StatusCode != 200 {
		t.Errorf("error: %v", rsp.StatusCode)
	}

	cancel()
	if err := eg.Wait(); err != nil {
		t.Errorf("error: %v", err)
	}
}
