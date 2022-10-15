//go:build wireinject
// +build wireinject

package main

import (
	"weeee9/wire-example/model"
	"weeee9/wire-example/router"
	"weeee9/wire-example/service"

	"github.com/google/wire"
)

func InitializeApp() (*app, error) {
	wire.Build(
		model.NewEngine,
		model.NewUserRepository,
		service.NewUserService,
		router.NewUserHandler,
		router.NewRouter,
		NewApp,
	)
	return &app{}, nil
}
