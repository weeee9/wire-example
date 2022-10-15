package main

import "net/http"

type app struct {
	router http.Handler
}

func NewApp(router http.Handler) *app {
	return &app{
		router: router,
	}
}
