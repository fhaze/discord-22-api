package router

import (
	"discord-22-api/router/handler"
	"github.com/labstack/echo/v4"
)

func Assign(e *echo.Echo) {
	v1 := e.Group("/22/api/v1")

	v1.GET("/users", handler.GetUsers)
	v1.GET("/users/:id", handler.GetUser)
	v1.POST("/users", handler.PostUser)
	v1.PUT("/users/:id/sum-message-count", handler.PutSumMessageCount)
	v1.PUT("/users/:id/sum-command-count", handler.PutSumCommandCount)
}
