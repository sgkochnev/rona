package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sgkochnev/rona/config"
	"github.com/sgkochnev/rona/pkg/httpserver"
	"github.com/sgkochnev/rona/pkg/logger"

	"github.com/sgkochnev/rona/internal/controller/http/v1"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// TODO initialize repo

	// TODO initialize usecase

	handler := http.NewServeMux()

	v1.NewRouter(handler, l, nil)
	// TODO initialize controller

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run -httpServer.Notify: %w", err))
	}

	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
