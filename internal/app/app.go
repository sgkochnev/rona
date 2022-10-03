package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sgkochnev/rona/config"
	"github.com/sgkochnev/rona/pkg/httpserver"
	"github.com/sgkochnev/rona/pkg/logger"

	v1 "github.com/sgkochnev/rona/internal/controller/http/v1"
	"github.com/sgkochnev/rona/internal/repo"
	"github.com/sgkochnev/rona/internal/usecase"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	r, err := repo.NewStore(cfg)
	if err != nil {
		l.Error("error: initialization repository failed: %v", err)
	}

	uc := usecase.NewManager(r, []byte(cfg.Secret.SignedKey))
	// handler := http.NewServeMux()

	handler := v1.NewRouter(l, uc)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	if err := httpServer.Shutdown(); err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
