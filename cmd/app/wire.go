//go:build wireinject
// +build wireinject

package main

import (
	"weeee9/wire-example/config"
	otelxorm "weeee9/wire-example/middleware/otel-xorm"
	"weeee9/wire-example/model"
	"weeee9/wire-example/router"

	"github.com/google/wire"
)

func InitializeApp(cfg config.Config) (*app, error) {
	wire.Build(
		otelxorm.NewTracingHook,
		model.NewEngine,
		model.NewUserRepository,
		router.NewUserHandler,
		router.NewRouter,
		NewApp,
	)
	return &app{}, nil
}
