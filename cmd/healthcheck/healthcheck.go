package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/geekcamp-vol11-team30/backend/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	url := &url.URL{}
	url.Scheme = "http"
	url.Host = net.JoinHostPort("localhost", fmt.Sprintf("%d", cfg.Port))
	url.Path = "/health"

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(context, "GET", url.String(), nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	rsp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	if rsp.StatusCode != 200 {
		panic(fmt.Sprintf("status code is not 200: %d", rsp.StatusCode))
	}

}
