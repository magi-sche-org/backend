package main

import (
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	cfg, err := config.New()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	db, err := db.NewDB(cfg, logger)
	if err != nil {
		logger.Fatal("failed to connect db", zap.Error(err))
	}
	defer db.Close()

	err = goose.SetDialect("mysql")
	if err != nil {
		logger.Fatal("failed to set dialect", zap.Error(err))
	}

	err = goose.Up(db, "db/migrations")
	if err != nil {
		logger.Fatal("failed to migrate", zap.Error(err))
	}
}
