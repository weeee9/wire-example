package main

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	app, err := InitializeApp()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init app")
	}

	srv := &http.Server{
		Addr:              ":8080",
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
