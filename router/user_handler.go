package router

import (
	"net/http"
	"strconv"
	"weeee9/wire-example/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (h UserHandler) getAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"users": users,
	})
}

func (h UserHandler) getUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	user, err := h.service.GetUserByID(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"users": user,
	})
}
