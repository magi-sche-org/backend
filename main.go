package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/geekcamp-vol11-team30/backend/handler"
	"github.com/geekcamp-vol11-team30/backend/pb"
	"github.com/geekcamp-vol11-team30/backend/store"
)

var db *sql.DB

func main() {
	log.Println("Go app started")
	if err := run(context.Background()); err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	// defer stop()

	db, err := store.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterEventServer(s, handler.NewEventServer(db))
	pb.RegisterAuthorizeServer(s, handler.NewAuthorizationServer(db))

	err = db.Ping()
	if err != nil {
		log.Println("error on db.Ping()")
		return err
	}
	log.Println("db.Ping() success")
	reflection.Register(s)
	go func() {
		log.Printf("start server on port %d", 50051)
		s.Serve(l)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		log.Printf("start health check server on port %d", 8080)
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
		http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	s.GracefulStop()
	return nil
}
