package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	config := mysql.NewConfig()

	config.Net = "tcp"
	// config.Addr = cfg.DBHost + ":" + string(cfg.DBPort)
	config.Addr = fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort)
	config.User = cfg.DBUser
	config.Passwd = cfg.DBPassword
	config.DBName = cfg.DBName
	config.ParseTime = true
	dsn := config.FormatDSN()
	log.Println(dsn)
	return sql.Open("mysql", dsn)
}
