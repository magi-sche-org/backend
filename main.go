package main

import (
	"context"
	"database/sql"
	"log"
	"os"

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
	db, err := store.ConnectDB()
	if err != nil {
		return err
	}

	log.Println("Hello, world!")
	log.Println(db)
	err = db.Ping()
	if err != nil {
		log.Println("error on db.Ping(): ", err)
		return err
	}
	log.Println("db.Ping() success")
	return nil
}
