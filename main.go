package main

import (
	"discord-22-api/config"
	"discord-22-api/router"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("X-22-KEY")
			if key == config.Instance().RootApiKey {
				return next(c)
			}
			return c.NoContent(http.StatusUnauthorized)
		}
	})
	router.Assign(e)
	e.Logger.Fatal(e.Start(":8888"))
}
