package main

import (
	"context"
	"fmt"
	"net"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/controller"
	"github.com/geekcamp-vol11-team30/backend/db"
	"github.com/geekcamp-vol11-team30/backend/middleware"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/router"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/geekcamp-vol11-team30/backend/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := run(context.Background(), logger); err != nil {
		sugar.Fatal(err)
	}
}

func run(ctx context.Context, logger *zap.Logger) error {
	logger.Info("magische starting...")

	cfg, err := config.New()
	if err != nil {
		return err
	}
	db, err := db.NewDB(cfg, logger)
	if err != nil {
		return err
	}
	boil.SetDB(db)
	boil.DebugMode = cfg.Env == "dev"

	ur := repository.NewUserRepository(db)
	ar := repository.NewAuthRepository(db)
	er := repository.NewEventRepository(db)
	uv := validator.NewUserValidator()
	uu := usecase.NewUserUsecase(ur, uv)
	au := usecase.NewAuthUsecase(cfg, logger, ar)
	eu := usecase.NewEventUsecase(er)

	em := middleware.NewErrorMiddleware(logger, uu)
	atm := middleware.NewAccessTimeMiddleware()
	am := middleware.NewAuthMiddleware(cfg, logger, au, uu)

	uc := controller.NewUserController(uu)
	ac := controller.NewAuthController(cfg, uu, au)
	ec := controller.NewEventController(eu)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		logger.Fatal("failed to listen port", zap.Error(err))
	}

	e := router.NewRouter(cfg, logger, em, atm, am, uc, ac, ec)
	s := NewServer(e, l, logger)
	return s.Run(ctx)
}
