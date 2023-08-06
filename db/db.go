package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/geekcamp-vol11-team30/backend/config"
	"go.uber.org/zap"
)

func NewDB(cfg *config.Config, logger *zap.Logger) (*sql.DB, error) {
	// config := mysql.NewConfig()
	// config.Net = "tcp"
	// config.Addr = fmt.Sprintf("%s:%d", cfg.MySQL.Host, cfg.MySQL.Port)
	// config.User = cfg.MySQL.User
	// config.Passwd = cfg.MySQL.Password
	// config.DBName = cfg.MySQL.DBName
	// config.ParseTime = true
	// config.Params = map[string]string{
	// 	"charset": "utf8mb4",
	// }
	// dsn := config.FormatDSN()
	db, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			// "%s:%s@%s:%d/%s?parseTime=true",
			cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host,
			cfg.MySQL.Port, cfg.MySQL.DBName,
		),
	)

	// logger.Info(dsn)
	// logger.Info(fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%d)/%s?parseTime=true",
	// 	cfg.MySQL.Host, cfg.MySQL.User, cfg.MySQL.Host,
	// 	cfg.MySQL.Port, cfg.MySQL.DBName,
	// ))
	if err != nil {
		logger.Error("failed to open db", zap.Error(err))
		return nil, err
	}

	ctx, canncel := context.WithTimeout(context.Background(), 10*time.Second)
	defer canncel()
	if err := db.PingContext(ctx); err != nil {
		logger.Error("failed to ping db", zap.Error(err))
		return nil, err
	}
	return db, nil
}
