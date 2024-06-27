package internal

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"os"
	"os/signal"
	"syscall"
	"testTask/cmd/internal/config"
	"testTask/cmd/internal/db/mysql"
	"testTask/cmd/internal/server"
	"testTask/cmd/internal/services"
)

type App struct {
	ctx context.Context
	db  io.Closer
	srv server.IServer
}

func (a *App) Close() {
	var err error

	if err = a.srv.Close(); err != nil {
		log.Error("Error while shutdown server: ", err)
	}

	if err = a.db.Close(); err != nil {
		log.Error("Error while close database connection: ", err)
	}
}

func (a *App) Run() {
	go func() {
		if err := a.srv.Start(); err != nil {
			log.Error(err)
		}
	}()

	done := make(chan os.Signal)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	<-done

	a.Close()
}

func NewApp(cfg config.Config) (*App, error) {
	d, err := mysql.NewDatabase(cfg.DatabaseConfig.DSN, cfg.DatabaseConfig.NeedRecreateDB)
	if err != nil {
		return nil, err
	}

	srv := server.NewServer(cfg.ServerConfig)

	authService := services.NewAuthService(d)
	userService := services.NewUserProfileService(d)

	srv.RegisterMiddleware(authService)
	srv.InitRoutes(userService)

	return &App{
		srv: srv,
		db:  d,
		ctx: context.Background(),
	}, nil
}
