package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler UserHandler) http.Handler {
	r := gin.Default()

	r.GET("/v1/users", userHandler.getAllUsers)
	r.GET("/v1/users/:id", userHandler.getUser)

	return r
}
