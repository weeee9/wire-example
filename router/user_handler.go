package router

import (
	"net/http"
	"strconv"
	"weeee9/wire-example/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo model.UserRepository
}

func NewUserHandler(repo model.UserRepository) UserHandler {
	return UserHandler{
		repo: repo,
	}
}

func (h UserHandler) getAllUsers(c *gin.Context) {
	users, err := h.repo.GetAllUsers(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
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
			"message": err.Error(),
		})
		return
	}

	user, err := h.repo.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"users": user,
	})
}
