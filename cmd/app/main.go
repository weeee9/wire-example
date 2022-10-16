package main

import (
	"net/http"
	"time"

	"weeee9/wire-example/config"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Environ()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init config")
	}

	app, err := InitializeApp(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init app")
	}

	srv := &http.Server{
		Addr:              cfg.Server.Port,
		Handler:           app.router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("server error")
	}
}
