package main

import (
	"log"

	"github.com/sgkochnev/rona/config"
	"github.com/sgkochnev/rona/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error %s", err)
	}

	app.Run(cfg)
}
