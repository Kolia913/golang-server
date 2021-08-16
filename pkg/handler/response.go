package handler

import (
	"github.com/Kolia913/todo-app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}
